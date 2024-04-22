package main

import (
	"auto-daily-report/src/config"
	"auto-daily-report/src/router"
	"auto-daily-report/src/utils"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

func main() {
	logger.Init("Logger", true, false, io.Discard)
	readConfig, err := utils.ReadConfigFromFile()
	if err != nil {
		log.Fatal(err)
	}

	if readConfig.Run.Mode == config.ModeDevelopment {
		logger.Info("Running in debug mode")
		gin.SetMode(gin.DebugMode)
	}

	if readConfig.Run.Mode == config.ModeProduction {
		logger.Info("Running in release mode")
		gin.SetMode(gin.ReleaseMode)
	}

	route := router.Router(*readConfig)

	err = route.Run()
	if err != nil {
		log.Fatal(err)
	}
}