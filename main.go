package main

import (
	"fmt"
	"tabby-sync/configs"
	_ "tabby-sync/configs"
	"tabby-sync/oauth"
	"tabby-sync/oauth/github"
	"tabby-sync/routes"
	_ "tabby-sync/services"
)

func main() {
	router := routes.Init()

	oauth.Register(router, github.New())
	_ = router.Run(fmt.Sprintf(":%d", configs.Platform.Port))
}
