package routes

import (
	"github.com/gin-gonic/gin"
	"tabby-sync/apis"
	"tabby-sync/middlewares"
)

func Init() *gin.Engine {
	r := gin.Default()

	api := r.Group("api/1")
	{

		user := api.Group("user")
		{
			user.Use(middlewares.GetUserID)
			user.GET("", apis.GetUser)
		}

		config := api.Group("configs")
		{
			config.Use(middlewares.GetUserID)
			config.GET("", apis.GetConfigList)
			config.GET(":id", apis.GetConfig)

			config.POST("", apis.CreateConfig)

			config.PATCH(":id", apis.UpdateConfig)

			config.DELETE(":id", apis.DeleteConfig)
		}
	}

	return r
}
