package router

import (
	"auto-daily-report/src/config/constants/environments"
	"auto-daily-report/src/middlewares"
	"auto-daily-report/src/wire"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/logger"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func Router() *gin.Engine {
	// setup telegram bot
	bot, err := tgbotapi.NewBotAPI(environments.TelegramBotApiKey)
	if err != nil {
		logger.Fatal(err)
	}

	client := openai.NewClient(environments.OpenAIApiKey)

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "x-api-key"}
	router.Use(cors.New(corsConfig))

	telegramController := wire.GetTelegramController(bot, client)
	githubController := wire.GetGitHubController()

	apiRoute := router.Group("/api")
	apiRoute.Use()
	{
		apiRoute.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
		telegramController.RegisterRoutes(apiRoute.Group("/telegram"))

		githubRoute := apiRoute.Group("/github")
		githubRoute.Use(middlewares.GitHubMiddleware())
		{
			githubController.RegisterRoutes(githubRoute)
		}

	}
	return router
}
