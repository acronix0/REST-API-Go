package v1

import (
	"github.com/acronix0/REST-API-Go/internal/service"
	"github.com/acronix0/REST-API-Go/pkg/auth"
	"github.com/gin-gonic/gin"
)
type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initCategoriesRoutes(v1)
		h.initUsersRoutes(v1)
		h.initAuthsRoute(v1)
	}
}