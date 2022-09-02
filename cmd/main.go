package main

import (
	"github.com/neodejack/CaptainJack/hex/adapters/application/api"
	"github.com/neodejack/CaptainJack/hex/adapters/core/echoing"
	"github.com/neodejack/CaptainJack/hex/adapters/frameworks/driving/telegramBotUi"
)

func main() {
	core := echoing.New()

	applicationAPI := api.NewApplication(core)

	teleBotAdapter := telegramBotUi.NewAdapter(applicationAPI)

	teleBotAdapter.Run()
}
