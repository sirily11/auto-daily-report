//go:build wireinject
// +build wireinject

// / This file is used to generate wire.go file using wire tool
// / Run `go generate ./...` to generate wire.go file
package wire

import (
	"auto-daily-report/src/controllers/github"
	tg "auto-daily-report/src/controllers/telegram"
	github3 "auto-daily-report/src/repositories/github"
	github2 "auto-daily-report/src/services/github"
	"auto-daily-report/src/services/report"
	telegram2 "auto-daily-report/src/services/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"
	"github.com/sashabaranov/go-openai"
)

func GetTelegramController(telegram *tgbotapi.BotAPI, client *openai.Client) *tg.Controller {
	wire.Build(tg.NewController, telegram2.NewCommandService, report.NewService, github3.NewIssueRepository)
	return &tg.Controller{}
}

func GetGitHubController() *github.Controller {
	wire.Build(github.NewController, github2.NewService, github3.NewIssueRepository)
	return &github.Controller{}
}
