package telegram

import (
	"auto-daily-report/src/pkgs/command"
	"auto-daily-report/src/services/report"
	"auto-daily-report/src/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandServiceInterface interface {
	ReplyToMessage(message tgbotapi.Message) (types.ServiceAction, error)
	RunCommand(message tgbotapi.CallbackQuery) (types.ServiceAction, error)
	AvailableCommands() []tgbotapi.BotCommand
}

type CommandService struct {
	bot           *tgbotapi.BotAPI
	commands      []command.ICommand
	reportService report.ServiceInterface
}

func NewCommandService(bot *tgbotapi.BotAPI, reportService report.ServiceInterface) CommandServiceInterface {
	return &CommandService{
		bot: bot,
		commands: []command.ICommand{
			command.NewReportCommand(bot, reportService),
		},
	}
}

func (c *CommandService) getKeyboardButton() []tgbotapi.InlineKeyboardButton {
	var buttons []tgbotapi.InlineKeyboardButton
	for _, c := range c.commands {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(c.Name(), c.Name()))
	}
	return buttons
}

func (c *CommandService) handleStart(chatId int64) error {
	inline := tgbotapi.NewInlineKeyboardRow(
		c.getKeyboardButton()...,
	)

	msg := tgbotapi.NewMessage(0, "*Welcome to Auto Daily Report Bot*\n\nPlease select the command below:")
	//msg.ReplyMarkup = keyboard
	msg.ChatID = chatId
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inline)
	msg.ParseMode = "MarkdownV2"

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommandService) ReplyToMessage(message tgbotapi.Message) (types.ServiceAction, error) {
	switch message.Command() {
	case "start", "help":
		err := c.handleStart(message.Chat.ID)
		if err != nil {
			return types.ActionStop, err
		}
		return types.ActionStop, nil

	default:
		return types.ActionContinue, nil
	}
}

func (c *CommandService) RunCommand(message tgbotapi.CallbackQuery) (types.ServiceAction, error) {
	for _, cmd := range c.commands {
		if cmd.Name() == message.Data {
			return cmd.Run(message)
		}
	}
	return types.ActionContinue, nil

}

func (c *CommandService) AvailableCommands() []tgbotapi.BotCommand {
	return []tgbotapi.BotCommand{
		{Command: "start", Description: "Start the bot"},
		{Command: "help", Description: "Show help"},
	}
}
