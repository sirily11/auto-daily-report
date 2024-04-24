package telegram

import (
	"auto-daily-report/src/config/constants/environments"
	telegram2 "auto-daily-report/src/services/telegram"
	"auto-daily-report/src/types"
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/logger"
	"net/http"
	url2 "net/url"
)

type ControllerInterface interface {
	Chat(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type Controller struct {
	telegram       *tgbotapi.BotAPI
	commandService telegram2.CommandServiceInterface
}

func NewController(telegram *tgbotapi.BotAPI, commandService telegram2.CommandServiceInterface) *Controller {
	cfg := tgbotapi.NewSetMyCommands(commandService.AvailableCommands()...)
	_, err := telegram.Request(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	return &Controller{
		telegram:       telegram,
		commandService: commandService,
	}
}

func (c *Controller) Chat(ctx *gin.Context) {
	var body tgbotapi.Update
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// regular message
	if body.Message != nil {
		result, err := c.commandService.ReplyToMessage(*body.Message)
		if err != nil {
			logger.Error(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if result == types.ActionStop {
			ctx.JSON(http.StatusOK, gin.H{"message": "Successfully received message from telegram"})
			return
		}
	}

	if body.CallbackQuery != nil {
		result, err := c.commandService.RunCommand(*body.CallbackQuery)
		if err != nil {
			logger.Error(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if result == types.ActionStop {
			ctx.JSON(http.StatusOK, gin.H{"message": "Successfully received message from telegram"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully received message from telegram"})
}

// Register is a function to register endpoint to telegram server
func (c *Controller) Register(ctx *gin.Context) {
	url, err := url2.JoinPath(environments.TelegramChatEndpoint, "/api/telegram/chat")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Infof("Registering to telegram server: %s", url)
	data, err := http.Get("https://api.telegram.org/bot" + c.telegram.Token + "/setWebhook?url=" + url)
	if err != nil {
		logger.Error(err)
		return
	}
	defer data.Body.Close()

	// Check if the response is successful
	if data.StatusCode != http.StatusOK {
		logger.Errorf("Failed to register to telegram server: %s", url)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to register to telegram server: %s", url)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully registered to telegram server: %s", url)})
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/chat", c.Chat)
	router.POST("/register", c.Register)
}
