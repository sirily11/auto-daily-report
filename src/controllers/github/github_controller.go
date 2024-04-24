package github

import (
	"auto-daily-report/src/dto/github"
	github2 "auto-daily-report/src/services/github"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/logger"
	"net/http"
)

type ControllerInterface interface {
	Webhook(ctx *gin.Context)
}

type Controller struct {
	telegram      *tgbotapi.BotAPI
	githubService github2.ServiceInterface
}

func NewController(githubService github2.ServiceInterface) *Controller {
	return &Controller{
		githubService: githubService,
	}
}

func (c *Controller) Webhook(ctx *gin.Context) {
	var body map[string]interface{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle closed GitHub issue
	if body["action"] == github.ClosedAction && body["issue"] != nil {
		issue := &github.IssueWebhookDTO{}
		err := json.Unmarshal(bytes, issue)
		if err != nil {
			logger.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.handleClosedIssue(issue)
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully received message from github"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully received message from github"})
}

func (c *Controller) handleClosedIssue(issue *github.IssueWebhookDTO) {
	issue.IsReported = false
	err := c.githubService.InsetIssue(*issue)
	if err != nil {
		logger.Error(err)

	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/webhook", c.Webhook)
}
