package command

import (
	"auto-daily-report/src/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ICommand interface {
	//Name returns the name of the command used to identify the command
	Name() string
	//Run executes the command
	Run(message tgbotapi.CallbackQuery) (types.ServiceAction, error)
}
