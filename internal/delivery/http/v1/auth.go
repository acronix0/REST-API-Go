package v1

import (
	"net/http"

	"github.com/acronix0/REST-API-Go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func (h *Handler)initAuthsRoute(api *gin.RouterGroup) {
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", h.register)
    authGroup.POST("/login", h.login)
		authGroup.PATCH("/password-reset", h.resetPassword)
		authGroup.POST("/token-refresh", h.refreshToken)
	}
	{
		authGroup.Use(h.userIdentity)
		authGroup.Use(h.authorizeRole(adminRole))
    authGroup.POST("/register-admin", h.registerAdmin)
	}
}

// @Summary User Registration
// @Description Register a new user with client role
// @Tags auth
// @Accept json
// @Produce json
// @Param input body service.UserRegisterInput true "User Registration Data"
// @Success 200 {object} service.Tokens
// @Failure 400 {object} response "Invalid input or registration failed"
// @Router /auth/register [post]
func (h *Handler) register(c *gin.Context){
	var user service.UserRegisterInput
  if err := c.ShouldBindJSON(&user); err!= nil {
    newResponse(c, http.StatusBadRequest,"Invalid credentials")
    return
  }
	if err := validator.New().Struct(user); err != nil {
    validationErrors := err.(validator.ValidationErrors)
    c.JSON(http.StatusBadRequest, validationErrors.Error())
    return
  }
	userAgent := c.GetHeader("User-Agent")
  tokens,err := h.services.Users().SignUp(c.Request.Context(), user,userAgent, clientRole)
  if err!= nil {
    newResponse(c, http.StatusInternalServerError,"Register failed")
    return
  }
  c.JSON(http.StatusOK, tokens)
}

// @Summary Admin Registration
// @Description Register a new user with admin role
// @Tags auth
// @Accept json
// @Produce json
// @Param input body service.UserRegisterInput true "Admin Registration Data"
// @Success 200 {object} service.Tokens "Tokens for the authenticated admin"
// @Failure 400 {object} response "Invalid input or registration failed"
// @Router /auth/register-admin [post]
func (h *Handler) registerAdmin(c *gin.Context){
	var user service.UserRegisterInput
  if err := c.ShouldBindJSON(&user); err!= nil {
    newResponse(c, http.StatusBadRequest,"Invalid credentials")
    return
  }
	if err := validator.New().Struct(user); err != nil {
    validationErrors := err.(validator.ValidationErrors)
    c.JSON(http.StatusBadRequest, validationErrors.Error())
    return
  }
	userAgent := c.GetHeader("User-Agent")
  tokens,err := h.services.Users().SignUp(c.Request.Context(), user,userAgent, adminRole)
  if err!= nil {
    newResponse(c, http.StatusBadRequest,"Register failed")
    return
  }
  c.JSON(http.StatusOK, tokens)
}

// @Summary User Login
// @Description Authenticate user and get JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param input body service.UserLoginInput true "Login Data"
// @Success 200 {object} service.Tokens "Tokens for the authenticated user"
// @Failure 400 {object} response "Invalid credentials"
// @Failure 401 {object} response "Login failed"
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var user service.UserLoginInput
  if err := c.ShouldBindJSON(&user); err!= nil {
    newResponse(c, http.StatusBadRequest,"Invalid credentials")
    return
  }
	userAgent := c.GetHeader("User-Agent")
  tokens, err := h.services.Users().SignIn(c.Request.Context(), user, userAgent)
  if err!= nil {
    newResponse(c, http.StatusUnauthorized, "Login failed")
    return
  }
  c.JSON(http.StatusOK, tokens)
}

// @Summary Reset Password
// @Security APiKeyAuth
// @Description Reset password for the authenticated user
// @Tags auth
// @Accept json
// @Produce json
// @Param newPassword body string true "New password"
// @Success 200 {object} response "Password successfully reset"
// @Failure 400 {object} response "Invalid user ID or invalid credentials"
// @Router /auth/password-reset [patch]
func (h *Handler) resetPassword(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var newPassword string
  if err := c.ShouldBindJSON(&newPassword); err!= nil {
    newResponse(c, http.StatusBadRequest,"Invalid credentials")
    return
  }
	
	if err:= h.services.Users().ChangePassword(c.Request.Context(), id,newPassword);err != nil {
		newResponse(c, http.StatusBadRequest,"Reset password failed")
		return
	}
	 c.Status(http.StatusOK)
}
// @Summary Refresh Access Token
// @Tags auth
// @Description Refresh the access token using a valid refresh token
// @Accept json
// @Produce json
// @Param input body string true "Refresh Token"
// @Success 200 {object} service.Tokens "New access and refresh tokens"
// @Failure 400 {object} response "Invalid input"
// @Failure 401 {object} response "Invalid or expired refresh token"
// @Failure 500 {object} response "Failed to refresh tokens"
// @Router /auth/refresh [post]
func (h *Handler) refreshToken(c *gin.Context) {
	var refreshToken string

	if err := c.BindJSON(&refreshToken); err != nil {
	    newResponse(c, http.StatusBadRequest, "Invalid input")
	    return
	}
	if refreshToken == "" {
	    newResponse(c, http.StatusBadRequest, "Refresh token required")
	    return
	}
	userAgent := c.GetHeader("User-Agent")
			
	tokens, err := h.services.Users().RefreshTokens(c.Request.Context(), refreshToken,userAgent)
	if err != nil {
			newResponse(c, http.StatusInternalServerError, "Failed to refresh tokens")
	    return
	}

	c.JSON(http.StatusOK, gin.H{
	    "access_token":  tokens.AccessToken,
	    "refresh_token": tokens.RefreshToken, 
	})
	}