package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rhiskey/spotytg/auths"
	"github.com/rhiskey/spotytg/spotifydl"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rhiskey/spotytg/utils"
	"github.com/zmb3/spotify/v2"
	"log"
	"os"
)

var (
	ctx           context.Context
	spotifyClient *spotify.Client
	bot           *tgbotapi.BotAPI
	apiEntity     *structures.Api
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
	log.Printf("📢 Authorized on account %s", bot.Self.UserName)

	apiEntity = structures.NewApi(spotifyClient, bot)
}

func main() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		apiEntity.TelegramMessageConfig = msg

		if update.Message.IsCommand() {
			//	Extract the command from the Message.
			switch update.Message.Command() {
			case "start":
				utils.LogWithBot("🔗 Just send me a link in format https://open.spotify.com/track/111111111111?si=xxxxxxxxx", apiEntity)
			case "help":
				utils.LogWithBot("ℹ I understand /status, /send and /help.", apiEntity)
			case "status":
				utils.LogWithBot("\U0001F9EA Beta test", apiEntity)
			case "send":
				if len(update.Message.Entities) == 0 { // ignore any Message without Entities
					continue
				}

				cmds := update.Message.CommandArguments()
				fmt.Println(cmds)

				playlistURL := cmds

				//utils.LogWithBot("⏳ Please, wait...", apiEntity)

				savedFile, err := spotifydl.DonwloadFromURL(playlistURL, apiEntity, ctx)
				if err != nil {
					continue
				}

				file := tgbotapi.FilePath(savedFile)

				sendAudioRequest := tgbotapi.NewAudio(update.Message.Chat.ID, file)

				msg.Text = savedFile
				if _, err := bot.Send(sendAudioRequest); err != nil {
					log.Panic(err)
				}

				e := os.Remove(savedFile)
				if e != nil {
					log.Fatal(e)
				}
			default:
				utils.LogWithBot("😕 I dont know that command.", apiEntity)
			}
		}

		if len(update.Message.Entities) == 0 { // ignore any Message without Entities
			continue
		}

		//update.Message.CommandArguments()
		if !update.Message.Entities[0].IsURL() { // ignore any Message without URL Entity type
			continue
		}

		playlistURL := update.Message.Text

		//utils.LogWithBot("⏳ Please, wait...", apiEntity)

		savedFile, err := spotifydl.DonwloadFromURL(playlistURL, apiEntity, ctx)
		if err != nil {
			continue
		}

		file := tgbotapi.FilePath(savedFile)

		sendAudioRequest := tgbotapi.NewAudio(update.Message.Chat.ID, file)

		msg.Text = savedFile
		if _, err := bot.Send(sendAudioRequest); err != nil {
			log.Panic(err)
		}

		e := os.Remove(savedFile)
		if e != nil {
			log.Fatal(e)
		}
	}
}
