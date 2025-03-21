package oauth

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Auth interface {
	SetRoute(router *gin.Engine, middlewares ...gin.HandlerFunc) Auth
	Init(config []byte, baseURL string) Auth
	SetRegister(register func(oauthType uint8, oauthTypeString, username, email, token string, oauthID any) (userID uint64, err error)) Auth
	GetConfig() (config []byte, err error)
	GetOAuthName() (oauthName string)
	GetButtonHTML() (html string)
	GetError() error
}
type oauthManager struct {
	oauthPlugins map[string]*Auth
	sync.RWMutex
}
