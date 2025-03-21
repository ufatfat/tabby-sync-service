package oauth

import (
	"github.com/gin-gonic/gin"
	"tabby-sync/configs"
	"tabby-sync/middlewares"
	"tabby-sync/services"
)

var Manager *oauthManager

func init() {
	Manager = &oauthManager{
		oauthPlugins: make(map[string]*Auth),
	}
}

func Register(r *gin.Engine, auth Auth) {
	oauth := auth.Init([]byte(configs.OAuth[auth.GetOAuthName()].(string)), "/oauth").SetRoute(r, middlewares.GetUserID).SetRegister(services.NewUser)
	if err := oauth.GetError(); err != nil {
		panic(err)
	}
	oauthName := oauth.GetOAuthName()
	Manager.oauthPlugins[oauthName] = &oauth
}
