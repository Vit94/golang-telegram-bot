package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Vit94/golang-telegram-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	fmt.Println(tgClient)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token to access telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token not passed")
	}

	return *token
}
