package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/acronix0/REST-API-Go/internal/config"
	"github.com/acronix0/REST-API-Go/internal/database"
	delivery "github.com/acronix0/REST-API-Go/internal/delivery/http"
	"github.com/acronix0/REST-API-Go/internal/repository"
	server "github.com/acronix0/REST-API-Go/internal/server"
	"github.com/acronix0/REST-API-Go/internal/service"
	"github.com/acronix0/REST-API-Go/pkg/auth"
	"github.com/acronix0/REST-API-Go/pkg/hash"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/acronix0/REST-API-Go-protos/gen/go/auth"
)


// @title dm-trade API
// @version 1.0
// @description This is a sample server API.
// @termsOfService http://example.com/terms/

// @host localhost:8081
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Run(configPath string) {
	cfg := config.MustLoad(configPath)
	fmt.Println(cfg)
	log := setUpLogger(cfg.Env)

	postgreClient, err := database.New(&cfg.SqlConnection)
	if err != nil {
		log.Error(err.Error())
	}
	tokenManager, err := auth.NewManager(cfg.AuthConfig.JWT.Secret)
	if err!=nil {
		log.Error(err.Error())
	}
	hasher := hash.NewSHA1Hasher(cfg.AuthConfig.PasswordSalt)
	
	repos := repository.NewRepositories(postgreClient.GetDB())
	AuthServiceConn, err := grpc.NewClient("localhost:4404",grpc.WithTransportCredentials(insecure.NewCredentials()))
	authGrpcClient :=  pb.NewAuthClient(AuthServiceConn)
	if err != nil {
		log.Error(err.Error())
	}
	services, err := service.NewServices(service.Deps{
		Repos: repos,
		TokenManager: tokenManager,
		Hasher: hasher,
		AccessTokenTTL: cfg.AuthConfig.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.JWTConfig.RefreshTokenTTL,
		AuthClient: authGrpcClient,
	})
	if err != nil {
		log.Error(err.Error())
	}
	handlers := delivery.NewHandler(services, tokenManager)
	srv := server.NewServer(cfg,handlers.Init(cfg))
		go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Error("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Error("failed to stop server: %v", err)
	}

}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
