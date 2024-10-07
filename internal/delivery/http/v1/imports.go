package v1
import (
	//"net/http"

	"github.com/gin-gonic/gin"
)
func (h *Handler) InitImportsRoutes(api *gin.RouterGroup){
	importsGroup := api.Group("imports")
	{
		importsGroup.Use(h.userIdentity)
		importsGroup.Use(h.authorizeRole(adminRole))
	}
}