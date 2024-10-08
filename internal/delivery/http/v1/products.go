package v1

import (
	"net/http"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup){
	prodictsGroup := api.Group("/products")
  {
      prodictsGroup.GET("/", h.getProducts)
      prodictsGroup.POST("/search", h.searchProducts)
  }
}
// @Summary Get Products
// @Tags products
// @Description Get Products
// @Produce  json
// @Success 200 {array} []Product
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /products [get]
func(h *Handler) getProducts(c *gin.Context){
	products, err := h.services.Products().GetProducts(c.Request.Context())
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to get products")
    return
  }

  c.JSON(http.StatusOK, products)
}
// @Summary Search Products
// @Description Search for products based on various filters and pagination
// @Tags products
// @Accept json
// @Produce json
// @Param input body GetProductsQuery true "Search and filter parameters"
// @Success 200 {array} Product
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Router /products/search [post]
func (h *Handler) searchProducts(c *gin.Context){
  var input domain.GetProductsQuery
  if err := c.BindJSON(&input); err!= nil {
    newResponse(c, http.StatusBadRequest, "Invalid input")
    return
  }
  products, err := h.services.Products().GetByCredentials(c.Request.Context(),input)
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to search products")
    return
  }
  c.JSON(http.StatusOK, products)
}