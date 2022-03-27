package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rhiskey/spotytg/structures"
	"log"
)

func LogWithBot(message string, api *structures.Api) {
	//fmt.Println(message)
	//log.Println(message)
	//rollbar.Info(message)
	api.TelegramMessageConfig.Text = message
	if _, err := api.TelegramBot.Send(api.TelegramMessageConfig); err != nil {
		log.Panic(err)
	}
}

func SendMessage(msg tgbotapi.MessageConfig, api *structures.Api) {
	if _, err := api.TelegramBot.Send(msg); err != nil {
		log.Panic(err)
	}
}
