package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetConfig(configPath string) {
	if configPath == "" {
		configPath = "./config.yaml"
	}

	if len(os.Args) > 1 {
		arg := os.Args[1]
		viper.SetConfigFile(arg)
	} else {
		viper.SetConfigFile(configPath)
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}
