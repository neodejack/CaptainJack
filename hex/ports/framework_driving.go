package ports

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelePort interface {
	InlineKeyboardWithCADates(update tgbotapi.Update) tgbotapi.MessageConfig
	CACallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.MessageConfig
	AnnoyingEchor(update tgbotapi.Update) (tgbotapi.MessageConfig, error)
	Run()
}
