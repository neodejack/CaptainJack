package telebotui

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/neodejack/CaptainJack/hex/ports"
)

type Adapter struct {
	api ports.APIPorts
	secrets ports.SecretsPorts
}

func NewAdapter(api ports.APIPorts, secrets ports.SecretsPorts) *Adapter {
	return &Adapter{api: api, secrets: secrets}
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
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			echoedWords, err := telebot.api.EchoWords(update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, echoedWords)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err = bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}

		}
	}
}
