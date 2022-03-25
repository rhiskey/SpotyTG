package utils

import (
	"fmt"
	"github.com/bogem/id3v2"
	"github.com/rhiskey/spotytg/structures"
	"github.com/zmb3/spotify/v2"
	"log"
	"strings"
)

// TagFileWithSpotifyMetadata takes in a filename as a string and spotify metadata and uses it to tag the music
func TagFileWithSpotifyMetadata(fileName string, trackData spotify.SimpleTrack, api *structures.Api) string {
	var trackArtist []string
	for _, Artist := range trackData.Artists {
		trackArtist = append(trackArtist, Artist.Name)
	}
	artistTag := strings.Join(trackArtist[:], ",")

	mp3File, err := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if err != nil {
		panic(err)
	}
	defer func(mp3File *id3v2.Tag) {
		err := mp3File.Close()
		if err != nil {

		}
	}(mp3File)

	mp3File.SetTitle(trackData.Name)
	mp3File.SetArtist(artistTag)

	if err = mp3File.Save(); err != nil {
		LogWithBot(fmt.Sprintf("⛔ Error while saving a tag: ", err), api)
		log.Fatal("Error while saving a tag: ", err)
	}

	return fileName

}
