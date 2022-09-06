package main

import (
	"github.com/neodejack/CaptainJack/hex/adapters/application/api"
	"github.com/neodejack/CaptainJack/hex/adapters/core/echoing"
	"github.com/neodejack/CaptainJack/hex/adapters/frameworks/driving/telegramBotUi"
	"github.com/neodejack/CaptainJack/hex/ports"
)

func main() {
	var core ports.EchoingPort
	var appAdapter ports.APIPorts
	var teleBotAdapter ports.TeleAdapter

	core = echoing.NewAdapter()
	appAdapter = api.NewApplication(core)
	teleBotAdapter = telegramBotUi.NewAdapter(appAdapter)

	teleBotAdapter.Run()
}
