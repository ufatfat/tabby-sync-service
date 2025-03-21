package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tabby-sync/models"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetUser{Name: c.GetString("username")})
}
