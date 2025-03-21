package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var configs Configs
var Database *database
var Platform *platform
var OAuth oauth
var Admin *user

func init() {
	configPath := "./config.yaml"
	if os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	}
	content, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(content, &configs); err != nil {
		panic(err)
	}
	log.Printf("Loaded configs: %+v", configs)
	Database = &configs.Database
	Platform = &configs.Platform
	OAuth = configs.OAuth
	Admin = &configs.Admin
}
