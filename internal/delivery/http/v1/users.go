package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/acronix0/REST-API-Go/internal/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	
	userGroup := api.Group("/users")
	{
		userGroup.Use(h.userIdentity)
		userGroup.Use(h.authorizeRole(clientRole, adminRole))
		userGroup.GET("/profile", h.getUserProfile)
	}
	{
		userGroup.Use(h.userIdentity)
		userGroup.Use(h.authorizeRole(adminRole))
		userGroup.GET("/", h.getUsers)
		userGroup.PATCH("/{id}/block", h.blockUser)
		userGroup.PATCH("/{id}/unblock",h.unblockUser)
		userGroup.PATCH("/",h.updateProfile)
		userGroup.PATCH("/password",h.updatePassword)
	}

}
// @Summary Get User Profile
// @Security UsersAuth
// @Tags users
// @Description Get profile of authenticated user
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /users/profile [get]
func (h *Handler) getUserProfile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, "Invalid user ID")
		return
	}

	user, err := h.services.Users.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			newResponse(c, http.StatusNotFound, "User not found")	
		}else {
  		newResponse(c, http.StatusInternalServerError, "Failed to get user")
  	}
		return
	}
	
	c.JSON(http.StatusOK, user)
}
// @Summary Get Users Profiles
// @Security UsersAuth
// @Tags users
// @Description Get profiles of users
// @Produce  json
// @Success 200 {array} []User
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /users [get]
func(h *Handler) getUsers(c *gin.Context){
	users, err := h.services.Users.GetUsers(c.Request.Context())
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to get users")
    return
  }

  c.JSON(http.StatusOK, users)
}
// @Summary Block User
// @Security UsersAuth
// @Tags users
// @Description Block User
// @Param id path int true "User ID"
// @Success 200 {string} string "OK"
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /{id}/block [patch]
func(h *Handler) blockUser(c *gin.Context){
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
  if err != nil {
    newResponse(c, http.StatusBadRequest, "Invalid user ID")
    return
  }

	err = h.services.Users.Block(c.Request.Context(), userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to unblock user")
		return
	}
	
  c.Status(http.StatusOK)
}
// @Summary Unblock user
// @Security UsersAuth
// @Tags users
// @Description Unblock User
// @Param id path int true "User ID"
// @Success 200 {string} string "OK"
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /{id}/unblock [patch]
func(h *Handler) unblockUser(c *gin.Context){
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
  if err != nil {
    newResponse(c, http.StatusBadRequest, "Invalid user ID")
    return
  }
	err = h.services.Users.Unblock(c.Request.Context(), userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to unblock user")
		return
	}
	
  c.Status(http.StatusOK)
}
// @Summary Update User Data
// @Security UsersAuth
// @Tags users
// @Description Update user profile information (excluding password)
// @Param id body int true "User ID"
// @Param userInput body updateUserInput true "User data"
// @Success 200 {string} string "OK"
// @Failure 400 {object} response
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /users [patch]
func (h *Handler) updateProfile(c *gin.Context){
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, "Invalid user ID")
		return
	}

 	var input service.UpdateUserInput
  if err := c.BindJSON(&input); err != nil {
    newResponse(c, http.StatusBadRequest, "Invalid input")
    return
  }

	if err := h.services.Users.UpdateProfile(c.Request.Context(), id, input); err!=nil{
		newResponse(c, http.StatusInternalServerError, "Failed to update user")
    return
	}
	c.Status(http.StatusOK)
}

// @Summary Change user password
// @Security UsersAuth
// @Tags users
// @Description change user password
// @Param id path int true "User ID"
// @Param input body  true "New password"
// @Success 200 {string} string "OK"
// @Failure 400 {object} response
// @Failure 401 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /users/password [patch]
func (h *Handler) updatePassword(c *gin.Context){
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, "Invalid user ID")
		return
	}

 	var newPass string
  if err := c.BindJSON(&newPass); err != nil {
    newResponse(c, http.StatusBadRequest, "Invalid input")
    return
  }

	if err := h.services.Users.ChangePassword(c.Request.Context(), id, newPass); err!=nil{
		newResponse(c, http.StatusInternalServerError, "Failed to change password")
    return
	}
	c.Status(http.StatusOK)
}