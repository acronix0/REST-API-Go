package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initCategoriesRoutes(api *gin.RouterGroup){
	categoriesGroup := api.Group("/categories")
  {
      categoriesGroup.GET("/", h.getCategories)
  }
}
// @Summary Get categories
// @Tags categories
// @Description Get categories
// @Produce  json
// @Success 200 {array} []Category
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /users [get]
func(h *Handler) getCategories(c *gin.Context){
	categories, err := h.services.Categories().GetCategories(c.Request.Context())
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to get categories")
    return
  }

  c.JSON(http.StatusOK, categories)
}