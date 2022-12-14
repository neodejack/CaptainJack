package secrets

import (
	"errors"
	"log"
	"os"
)

type AdapterOS struct {
}

func NewAdapterOS() *AdapterOS {
	return &AdapterOS{}
}

func (secrets AdapterOS) GetTeleToken() string {
	token := os.Getenv("TOKEN")
	if token == "" {
		err := errors.New("Couldn't initialize the bot with the token")
		log.Panic(err)
	}
	return token

}
