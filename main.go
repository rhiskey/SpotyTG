package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rhiskey/spotytg/auths"
	"github.com/rhiskey/spotytg/spotifydl"
	"github.com/zmb3/spotify/v2"
	"log"
	"os"
)

var (
	ctx           context.Context
	spotifyClient *spotify.Client
	bot           *tgbotapi.BotAPI
)

func init() {
	spotifyClient = auths.AuthSpotifyWithCreds()
	ctx = context.Background()

	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func main() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if len(update.Message.Entities) == 0 { // ignore any Message without Entities
			continue
		}

		if !update.Message.Entities[0].IsURL() { // ignore any Message without URL Entity type
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		playlistURL := update.Message.Text

		msg.Text = "⏳ Downloading..."
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}

		savedFile := spotifydl.DonwloadFromURL(playlistURL, spotifyClient, ctx)

		file := tgbotapi.FilePath(savedFile)

		sendAudioRequest := tgbotapi.NewAudio(update.Message.Chat.ID, file)

		msg.Text = savedFile
		if _, err := bot.Send(sendAudioRequest); err != nil {
			panic(err)
		}

		e := os.Remove(savedFile)
		if e != nil {
			log.Fatal(e)
		}
	}
}
