package main

import (
	"fmt"
	"tabby-sync/configs"
	_ "tabby-sync/configs"
	"tabby-sync/routes"
	_ "tabby-sync/services"
)

func main() {
	router := routes.Init()
	_ = router.Run(fmt.Sprintf(":%d", configs.Port))
}
