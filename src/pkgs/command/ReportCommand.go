package command

import (
	"auto-daily-report/src/services/report"
	"auto-daily-report/src/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ReportCommand struct {
	bot           *tgbotapi.BotAPI
	reportService report.ServiceInterface
}

func NewReportCommand(bot *tgbotapi.BotAPI, reportService report.ServiceInterface) ICommand {
	return &ReportCommand{
		bot:           bot,
		reportService: reportService,
	}
}

func (c *ReportCommand) Name() string {
	return "📔 Generate report"
}

func (c *ReportCommand) Run(message tgbotapi.CallbackQuery) (types.ServiceAction, error) {
	generatedReport, err := c.reportService.GenerateReport()
	if err != nil {
		replyMessage := tgbotapi.NewMessage(message.Message.Chat.ID, fmt.Sprintf("Failed to generate report: %s", err.Error()))
		_, err = c.bot.Send(replyMessage)
		return types.ActionContinue, err
	}
	replyMessage := tgbotapi.NewMessage(message.Message.Chat.ID, generatedReport)
	replyMessage.ParseMode = "Markdown"
	_, err = c.bot.Send(replyMessage)
	if err != nil {
		return types.ActionContinue, err
	}
	return types.ActionContinue, nil
}
