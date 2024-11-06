package v1

import (
	"net/http"

	"github.com/acronix0/REST-API-Go/internal/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initOrdersRoutes(api *gin.RouterGroup){
	ordersGroup := api.Group("/orders")
  {		ordersGroup.Use(h.userIdentity)
			ordersGroup.Use(h.authorizeRole(adminRole, clientRole))
    	ordersGroup.GET("/", h.getUserOrders)
      ordersGroup.POST("/", h.createOrder)
  }
}
// @Summary Get user orders
// @Tags orders
// @Description Get categories
// @Produce  json
// @Success 200 {array} []domain.Order
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /orders [get]
func(h *Handler) getUserOrders(c *gin.Context){
	userID, err := getUserId(c)
	 if err!= nil {
    newResponse(c, http.StatusUnauthorized, "Invalid user ID")
    return
  }
	orders, err := h.services.Orders().GetByUserId(c.Request.Context(), userID)
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to get user orders")
    return
  }

  c.JSON(http.StatusOK, orders)
}

// @Summary Create order
// @Tags orders
// @Description create order
// @Accept  json
// @Produce  json
// @Param input body service.CreateOrderInput true "User Registration Data"
// @Success 200 {string} string "OK"
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /orders [post]
func(h *Handler) createOrder(c *gin.Context){
	userID, err := getUserId(c)
	if err!= nil {
    newResponse(c, http.StatusUnauthorized, "Invalid user ID")
    return
  }
	var input service.CreateOrderInput
  if err := c.BindJSON(&input); err != nil {
    newResponse(c, http.StatusBadRequest, "Invalid input")
    return
  }
	input.UserID = userID
	err = h.services.Orders().Create(c.Request.Context(),input)
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to get user orders")
    return
  }

  c.Status(http.StatusOK)
}