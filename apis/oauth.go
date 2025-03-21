package apis

import "github.com/gin-gonic/gin"

func GetOAuthQueryParams(c *gin.Context) {}

func GitHubOAuth(c *gin.Context) {
	//code, state := c.Query("code"), c.Query("state")
}

func GitLabOAuth(c *gin.Context) {}
