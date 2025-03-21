package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

var configs Configs
var Database *database
var Port int

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
	Database = &configs.Database
	Port = configs.Port
}
