package github

import "time"

const (
	accessTokenURL  = "https://github.com/login/oauth/access_token"
	userAPIURL      = "https://api.github.com/user"
	oauthName       = "github"
	randomStatePath = oauthName + "/state"
	callbackPath    = oauthName + "/callback"
)

const (
	scope               = "read:user,user:email"
	stateExpireInterval = 3 * time.Minute
	cleanerInterval     = 30 * time.Minute
)
