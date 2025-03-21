package routes

import (
	"github.com/gin-gonic/gin"
	"tabby-sync/apis"
	"tabby-sync/middlewares"
)

func Init() *gin.Engine {
	r := gin.Default()

	api := r.Group("api")
	{
		api.Use(middlewares.GetUserID)

		api.GET("1/user", apis.GetUser)
		api.GET("1/configs", apis.GetConfigList)
		api.GET("1/configs/:id", apis.GetConfig)

		api.POST("1/configs", apis.CreateConfig)

		api.PATCH("1/configs/:id", apis.UpdateConfig)

		api.DELETE("1/configs/:id", apis.DeleteConfig)
	}

	return r
}
