package structures

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zmb3/spotify/v2"
)

type Api struct {
	SpotifyClient         *spotify.Client
	TelegramBot           *tgbotapi.BotAPI
	TelegramMessageConfig tgbotapi.MessageConfig
}

func NewApi(spotify *spotify.Client, telegramBot *tgbotapi.BotAPI) *Api {
	api := Api{
		SpotifyClient: spotify,
		TelegramBot:   telegramBot,
	}
	return &api
}
