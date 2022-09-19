package main

import (
	"github.com/neodejack/CaptainJack/hex/adapters/application/api"
	"github.com/neodejack/CaptainJack/hex/adapters/core/funstufffortina"
	"github.com/neodejack/CaptainJack/hex/adapters/frameworks/driven/secrets"
	"github.com/neodejack/CaptainJack/hex/adapters/frameworks/driving/telebotui"
	"github.com/neodejack/CaptainJack/hex/ports"
)

func main() {
	var core ports.EchoingPort
	var appAdapter ports.APIPorts
	var teleBotAdapter ports.TelePort
	var secretsAdapter ports.SecretsPorts

	//core = echoing.NewAdapter()
	//fun edition for tina
	core = funstufffortina.NewAdapter()
	appAdapter = api.NewApplication(core)
	secretsAdapter = secrets.NewAdapterOS()
	teleBotAdapter = telebotui.NewAdapter(appAdapter, secretsAdapter)

	teleBotAdapter.Run()
}
