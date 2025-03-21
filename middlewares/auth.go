package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tabby-sync/cache"
	"tabby-sync/services"
)

func GetUserID(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")[7:]

	// 先从cache拿userID
	userID := cache.GetUser(token)
	if userID == 0 {
		user, err := services.GetUser(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Set("username", user.Name)
		userID = user.UserID
		cache.SetUser(token, userID)
	}

	if userID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	c.Set("userID", userID)
}
