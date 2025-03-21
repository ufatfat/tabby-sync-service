package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tabby-sync/models"
	"tabby-sync/services"
)

func GetConfigList(c *gin.Context) {
	userID := c.GetUint64("userID")
	configs, err := services.GetConfigList(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, configs)
}

func GetConfig(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint64("userID")
	config, err := services.GetConfig(userID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, config)
}

func CreateConfig(c *gin.Context) {
	var req models.CreateConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID := c.GetUint64("userID")
	req.UserID = userID
	rst, err := services.CreateConfig(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rst)
}

func UpdateConfig(c *gin.Context) {
	var req models.UpdateConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	_id := c.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	userID := c.GetUint64("userID")
	rst, err := services.UpdateConfig(id, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rst)
}

func DeleteConfig(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint64("userID")
	if err = services.DeleteConfig(userID, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
