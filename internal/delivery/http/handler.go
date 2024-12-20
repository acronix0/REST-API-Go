package http

import (
	"fmt"
	"net/http"

	"github.com/acronix0/REST-API-Go/docs"
	"github.com/acronix0/REST-API-Go/internal/config"
	v1 "github.com/acronix0/REST-API-Go/internal/delivery/http/v1"
	"github.com/acronix0/REST-API-Go/internal/service"
	"github.com/acronix0/REST-API-Go/pkg/auth"
	"github.com/gin-gonic/gin"
	 "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

type Handler struct {
	services service.ServiceManager
	tokenManager auth.TokenManager
}
func NewHandler(service service.ServiceManager, tokenManager auth.TokenManager) *Handler{
	return &Handler{services: service, tokenManager: tokenManager}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTPConfig.Host, cfg.HTTPConfig.Port)
	if cfg.Env != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTPConfig.Host
	}

	if cfg.Env != config.EnvProd {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	} 

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}