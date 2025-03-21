package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	_oauth "tabby-sync/oauth"
)

var auth *oauth

func New() *oauth {
	return new(oauth)
}

func (a *oauth) Init(config []byte, basePath string) _oauth.Auth {
	a.basePath = basePath
	a.cache = new(cache)
	go a.cache.cleaner()

	var _err error
	if err := json.Unmarshal(config, &a.config); err == nil {
		return a
	} else {
		_err = errors.Join(_err, fmt.Errorf("parse config as json failed: %s", err.Error()))
	}
	if err := yaml.Unmarshal(config, &a.config); err == nil {
		return a
	} else {
		_err = errors.Join(_err, fmt.Errorf("parse config as yaml failed: %s", err.Error()))
	}
	a.err = errors.Join(a.err, _err, errors.New("cannot parse config"))
	return a
}

func (a *oauth) SetRoute(r *gin.Engine, middlewares ...gin.HandlerFunc) _oauth.Auth {
	auth := r.Group(a.basePath)
	{
		auth.Use(middlewares...)
		auth.GET(randomStatePath, func(c *gin.Context) {
			randString := genRandString(32, 32)
			a.cache.setState(randString)
			c.JSON(http.StatusOK, randString)
		})
		auth.GET(callbackPath, func(c *gin.Context) {
			code, state := c.Query("code"), c.Query("state")
			if !a.cache.stateAvailable(state) {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			if err := a.getUserInfo(code, state); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		})
	}
	return a
}

func (a *oauth) GetOAuthName() string {
	return "github"
}

func (a *oauth) GetButtonHTML() (html string) {
	// todo
	return ""
}

func (a *oauth) SetRegister(register func(oauthType uint8, oauthTypeString, username, email, token string, oauthID any) (userID uint64, err error)) _oauth.Auth {
	a.regFunc = register
	return a
}

func (a *oauth) GetError() error {
	return a.err
}

func (a *oauth) GetConfig() ([]byte, error) {
	// todo
	return nil, nil
}
