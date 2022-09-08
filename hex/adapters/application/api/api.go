package api

import (
	"github.com/neodejack/CaptainJack/hex/ports"
)

type Application struct {
	echor ports.EchoingPort
}

func NewApplication(core ports.EchoingPort) *Application {
	return &Application{core}
}

func (apia Application) EchoWords(words string) (string, error) {
	echoedWord, err := apia.echor.Echoing(words)
	if err != nil {
		return "error", err
	}
	return echoedWord, nil
}
