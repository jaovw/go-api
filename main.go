package main

import (
	"github.com/jaovw/go-api/config"
	"github.com/jaovw/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initialize()
}
