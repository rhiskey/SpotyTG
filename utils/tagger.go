package utils

import (
	"fmt"
	"github.com/bogem/id3v2"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rollbar/rollbar-go"
	"github.com/zmb3/spotify/v2"
	"log"
	"strconv"
	"strings"
	"time"
)

// TagFileWithSpotifyMetadata takes in a filename as a string and spotify metadata and uses it to tag the music
func TagFileWithSpotifyMetadata(fileName string, trackData spotify.FullTrack) string {

	albumTag := trackData.Album.Name
	trackArtist := []string{}
	for _, Artist := range trackData.Album.Artists {
		trackArtist = append(trackArtist, Artist.Name)
	}
	artistTag := strings.Join(trackArtist[:], ",")
	dateObject, _ := time.Parse("2006-01-02", trackData.Album.ReleaseDate)
	yearTag := dateObject.Year()
	albumArtImages := trackData.Album.Images

	mp3File, err := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if err != nil {
		rollbar.Error(err)
		panic(err)
	}
	defer func(mp3File *id3v2.Tag) {
		err := mp3File.Close()
		if err != nil {
			rollbar.Error(err)
			panic(err)
		}
	}(mp3File)

	mp3File.SetTitle(trackData.Name)
	mp3File.SetArtist(artistTag)
	mp3File.SetAlbum(albumTag)
	mp3File.SetYear(strconv.Itoa(yearTag))

	if len(albumArtImages) > 0 {
		albumArtURL := albumArtImages[0].URL
		albumArt, albumArtDownloadErr := DownloadFile(albumArtURL)
		if albumArtDownloadErr == nil {
			pic := id3v2.PictureFrame{
				Encoding:    id3v2.EncodingUTF8,
				MimeType:    "image/jpeg",
				PictureType: id3v2.PTFrontCover,
				Description: "Front cover",
				Picture:     albumArt,
			}
			mp3File.AddAttachedPicture(pic)
		} else {
			fmt.Println("An error occured while downloading album art ", err)
			rollbar.Error(err)
		}
	} else {
		fmt.Println("No album art found for ", trackData.Name)
		rollbar.Warning(err)
	}

	if err = mp3File.Save(); err != nil {
		rollbar.Critical(err)
		log.Fatal("Error while saving a tag: ", err)
	}

	return fileName
}

// TagFileWithSpotifyMetadataV2 takes in a filename as a string and spotify metadata and uses it to tag the music
func TagFileWithSpotifyMetadataV2(fileName string, trackData spotify.SimpleTrack, api *structures.Api) string {
	var trackArtist []string
	for _, Artist := range trackData.Artists {
		trackArtist = append(trackArtist, Artist.Name)
	}
	artistTag := strings.Join(trackArtist[:], ",")

	mp3File, err := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if err != nil {
		rollbar.Error(err)
		log.Panic(err)
	}
	defer func(mp3File *id3v2.Tag) {
		err := mp3File.Close()
		if err != nil {
			rollbar.Critical(err)
			log.Panic(err)
		}
	}(mp3File)

	mp3File.SetTitle(trackData.Name)
	mp3File.SetArtist(artistTag)

	if err = mp3File.Save(); err != nil {
		LogWithBot(fmt.Sprintf("⛔ Error while saving a tag: ", err), api)
		rollbar.Critical(err)
		log.Fatal("Error while saving a tag: ", err)
	}

	return fileName
}
