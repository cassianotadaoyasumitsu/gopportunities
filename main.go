package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

var (
	logger *config.Logger
)

func init() {
	err := config.LoadEnvVariables()
	if err != nil {
		return
	}
}

func main() {
	logger = config.GetLogger("main")
	// Initialize the config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
