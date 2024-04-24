package api

import (
	"auto-daily-report/src/config/constants/environments"
	"auto-daily-report/src/router"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	logger.Init("Logger", true, false, io.Discard)

	if environments.RunEnvironment != environments.ModeProduction {
		gin.SetMode(gin.DebugMode)
	}
	route := router.Router()

	route.ServeHTTP(w, r)
}
