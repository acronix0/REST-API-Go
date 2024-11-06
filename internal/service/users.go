package service

import (
	"context"
	"sync"
	"time"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/acronix0/REST-API-Go/internal/repository"
	"github.com/acronix0/REST-API-Go/pkg/auth"
	"github.com/acronix0/REST-API-Go/pkg/hash"
	"github.com/dgrijalva/jwt-go"
	pb "github.com/acronix0/REST-API-Go-protos/gen/go/auth"
)
type usersService struct {
	grpcClient pb.AuthClient
	userRepo repository.User
	authRepo repository.Auth
	tokenManager auth.TokenManager
	accessTokenTTL         time.Duration
	refreshTokenTTL        time.Duration
	hasher       hash.PasswordHasher
	roleCache   sync.Map  
}
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
func NewUsersService(authGrpcClient pb.AuthClient,hasher hash.PasswordHasher, userRepo repository.User, authRepo repository.Auth, accessTTL time.Duration, refreshTTL time.Duration) *usersService {
	return &usersService{
		grpcClient: authGrpcClient,
    hasher:       hasher,
    userRepo:  userRepo,
    authRepo:  authRepo,
    accessTokenTTL: accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

func (s *usersService) GetByID(ctx context.Context, id int) (domain.User, error){
	user, err := s.userRepo.GetById(ctx, id)
	if err == nil {
		return domain.User{}, err
	}
	
	return user, nil
}
func (s *usersService) GetUsers(ctx context.Context) ([]domain.User, error){
	users, err := s.userRepo.GetUsers(ctx)
	if err == nil {
		return nil, err
	}

	return users, nil
}
func (s *usersService) SignUp(ctx context.Context, input UserRegisterInput, deviceinfo string, role string) (Tokens, error) {

	hashedPassword, err := s.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}
	
	resp,err := s.grpcClient.SignUp(ctx, &pb.SignUpRequest{
		Email:    input.Email,
    Password: hashedPassword,
    DeviceInfo: deviceinfo, 
    Role: role,
		Username: input.Name,
		Phone: input.Phone,
  })
	if err != nil {
		return Tokens{}, err
	}
/* 	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
		Role: role,
		Blocked:  false,
	}

	err = s.userRepo.Create(ctx, &user)
	if err != nil {
		return Tokens{}, err
	} 

	return s.generateTokens(ctx, int(resp.UserId), deviceinfo)*/
	return Tokens{AccessToken: resp.GetJwtToken(), RefreshToken: resp.GetRefreshToken()}, nil
}
func (s *usersService) SignIn(ctx context.Context, input UserLoginInput, deviceInfo string) (Tokens, error) {


	hashedPassword, err := s.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}
	
	resp,err := s.grpcClient.SignIn(ctx, &pb.SignInRequest{
		Email:    input.Email,
    Password: hashedPassword,
    DeviceInfo: deviceInfo, 
  })


	/* user, err := s.userRepo.Login(ctx, input.Email, hashedPassword)
	if err != nil {
		return Tokens{}, err
	}
	return s.generateTokens(ctx, user.ID, deviceInfo) */
	return Tokens{ AccessToken: resp.GetJwtToken(), RefreshToken: resp.GetRefreshToken()}, nil
}

func (s *usersService) RefreshTokens(ctx context.Context, refreshToken string, deviceInfo string) (Tokens, error) {
	userID, err := s.tokenManager.Parse(refreshToken)
	if err != nil {
		return Tokens{}, err
	}
	valid, err := s.authRepo.ValidateRefreshToken(ctx, userID, refreshToken)
	if err != nil || !valid {
		return Tokens{}, err
	}

	return s.generateTokens(ctx, userID, deviceInfo)
}
func (s *usersService)generateTokens(ctx context.Context, userID int, deviceInfo string) (Tokens, error){
	accessToken, err := s.tokenManager.NewJWT(userID, time.Hour*1)
	if err != nil {
		return Tokens{}, err
	}
	newRefreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return Tokens{}, err
	}
	err = s.authRepo.DeleteRefreshToken(ctx, userID, deviceInfo)
	if err != nil {
		return Tokens{}, err
	}
	err = s.authRepo.SaveRefreshToken(ctx, userID, newRefreshToken, time.Now().Add(s.refreshTokenTTL), deviceInfo)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *usersService) ChangePassword(ctx context.Context, userId int, newPassword string) error {
	hashedPassword, err := s.hasher.Hash(newPassword)
	if err != nil {
		return err
	}

	return s.userRepo.ChangePassword(ctx, userId, hashedPassword)
}

func (s *usersService) Block(ctx context.Context, userID int) error {
	Blocked := true
	return s.userRepo.Update(ctx, repository.UpdateUserInput{
		ID: userID,
		Blocked: &Blocked,
	})
}
func (s *usersService) DeleteAllRefreshTokens(ctx context.Context, userID int) error{
	return s.authRepo.DeleteAllRefreshTokens(ctx, userID)	
}
func (s *usersService) Unblock(ctx context.Context, userID int) error {
	Blocked := false
	return s.userRepo.Update(ctx, repository.UpdateUserInput{
		ID: userID,
		Blocked: &Blocked,
	})
}

func (s *usersService) GetUserRole(ctx context.Context,userID int) (string, error) {
  if cachedRole, ok := s.roleCache.Load(userID); ok {
      return cachedRole.(string), nil
  }

  role, err := s.userRepo.GetRoleByUserID(ctx, userID)
  if err != nil {
      return "", err
  }
  s.roleCache.Store(userID, role)

  return role, nil
}

func (s *usersService) UpdateProfile(ctx context.Context, userID int, input UpdateUserInput) error {

  return s.userRepo.Update(ctx, repository.UpdateUserInput{
		ID: userID,
    Name: input.Name,
    Email: input.Email,
    Phone: input.Phone,
  })
	
}