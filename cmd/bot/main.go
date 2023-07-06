package main

import (
	"github.com/Dolaxome/instadownload-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5917885543:AAEAoOWqnR_jZcbagPmCfhftBn9r9xLPpAA")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
