package utils

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"strings"
)

// GetTrackNameFromUri is a function to find track in Spotify and return combined name
func GetTrackNameFromUri(trackUri string, spotifyClient *spotify.Client, ctx context.Context) (string, error) {
	var fullTrackName []string

	track, err := spotifyClient.GetTrack(ctx, spotify.ID(trackUri))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fullTrackName = append(fullTrackName, track.Name)

	for _, artist := range track.Artists {
		fullTrackName = append(fullTrackName, artist.Name)
	}

	return strings.Join(fullTrackName, " "), nil
}

// GetNameAndArtistsFromUri GetTrackNameFromUri is a function to find track in Spotify and return Name and Artists separately
func GetNameAndArtistsFromUri(trackUri string, spotifyClient *spotify.Client, ctx context.Context) (string, string, error) {
	var fullTrackName []string

	track, err := spotifyClient.GetTrack(ctx, spotify.ID(trackUri))
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	//fullTrackName = append(fullTrackName, track.Name)

	for _, artist := range track.Artists {
		fullTrackName = append(fullTrackName, artist.Name)
	}

	return track.Name, strings.Join(fullTrackName, " "), nil
}
