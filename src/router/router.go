package router

import (
	"auto-daily-report/src/config"
	"auto-daily-report/src/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(cfg config.Config) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "x-api-key"}
	router.Use(cors.New(corsConfig))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiRoute := router.Group("/api")
	apiRoute.Use(middlewares.APIKeyMiddleware())
	{

	}
	return router
}
