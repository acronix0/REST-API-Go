package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx           = "userID"
	clientRole = "Client"
	adminRole = "Admin"
)

func (h *Handler) authorizeRole(allowedRoles ...string) gin.HandlerFunc{
	return func(c *gin.Context){
		userID, ok := c.Get(userCtx)
		if !ok {
			newResponse(c, http.StatusUnauthorized, "userID not found in context")
		}
		userIDInt, ok := userID.(int)
		if !ok {
			newResponse(c, http.StatusUnauthorized, "userID is of invalid type")
      return
    }
		role, err := h.services.Users().GetUserRole(c.Request.Context(), userIDInt)
		if err != nil {
			newResponse(c, http.StatusForbidden, "forbidden")
      return
		}
	 	for _, allowedRole := range allowedRoles {
    	if role == allowedRole {
      	c.Next()
      	return
    	}
    }
		c.Next()
	}

}

func (h *Handler) userIdentity(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, id)
}
func getUserId(c *gin.Context) (int, error) {
	return getIdByContext(c, userCtx)
}


func (h *Handler) parseAuthHeader(c *gin.Context) (int, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return 0, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return 0, errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}

func getIdByContext(c *gin.Context, context string) (int, error) {
	idFromCtx, ok := c.Get(context)
	if !ok {
		return 0, errors.New("userCtx not found")
	}

	id, ok := idFromCtx.(int)
	if !ok {
		return 0, errors.New("userCtx is of invalid type")
	}
	return id, nil
}
