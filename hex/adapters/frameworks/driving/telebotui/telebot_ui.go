package telebotui

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/neodejack/CaptainJack/hex/ports"
)

type Adapter struct {
	api     ports.APIPorts
	secrets ports.SecretsPorts
}

func NewAdapter(api ports.APIPorts, secrets ports.SecretsPorts) *Adapter {
	return &Adapter{api: api, secrets: secrets}
}
func (telebot Adapter) InlineKeyboardForTina(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	courtDatesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("yes he's cute", "EWWWWWWW CUTE"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("no", "yeah he's just an average dude"),
		),
	)

	msg.ReplyMarkup = courtDatesKeyboard

	return msg

}

func (telebot Adapter) InlineKeyboardWithCADates(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	courtDatesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2022-09-20", "2022-09-20"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2022-09-22", "2022-09-22"),
		),
	)

	msg.ReplyMarkup = courtDatesKeyboard

	return msg

}

func (telebot Adapter) CACallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.MessageConfig {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}

	callbackWords := ""

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, callbackWords)

	return msg
}
func (telebot Adapter) AnnoyingEchor(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	echoedWords, err := telebot.api.EchoWords(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, echoedWords)
	msg.ReplyToMessageID = update.Message.MessageID
	return msg, err
}

func (telebot Adapter) Run() {
	token := telebot.secrets.GetTeleToken()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Couldn't initialize the bot with the token")
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			switch update.Message.Text {
			case "courtallocation":
				msg = telebot.InlineKeyboardWithCADates(update)
			case "zili?":
				msg = telebot.InlineKeyboardForTina(update)
			case "tina is here":
				msg, err = telebot.AnnoyingEchor(update)
				if err != nil {
					log.Panic(err)
				}
			}

		} else if update.CallbackQuery != nil {
			msg = telebot.CACallback(bot, update)
		}
		_, err = bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
