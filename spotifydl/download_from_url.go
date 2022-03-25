package spotifydl

import (
	"context"
	"errors"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rhiskey/spotytg/utils"
	"strings"
)

// DonwloadFromURL is a function to decide which type of content URL contains and download it
func DonwloadFromURL(spotifyURL string, api *structures.Api, ctx context.Context) (string, error) {
	var trackID string
	//var playlistID string
	//var albumID string

	var savedFile string

	if len(spotifyURL) == 0 {
		utils.LogWithBot("⚠ Spotify URL required.", api)
		return "", errors.New("spotify URL required")
	}

	splitURL := strings.Split(spotifyURL, "/")

	if len(splitURL) < 2 {
		utils.LogWithBot("⚠ Please enter the url copied from the spotify client.", api)
		return "", errors.New("wrong type of url")
	}

	spotifyID := splitURL[len(splitURL)-1]
	if strings.Contains(spotifyID, "?") {
		spotifyID = strings.Split(spotifyID, "?")[0]
	}

	//if strings.Contains(spotifyURL, "album") {
	//	albumID = spotifyID
	//	DownloadAlbum(albumID, api, ctx)
	//} else if strings.Contains(spotifyURL, "playlist") {
	//	playlistID = spotifyID
	//	DownloadPlaylist(playlistID, api, ctx)
	//} else
	if strings.Contains(spotifyURL, "track") {
		trackID = spotifyID
		var err error
		savedFile, err = DownloadSong(trackID, api, ctx)
		if err != nil {
			return "", err
		}
	} else {
		utils.LogWithBot("⚠ Only Spotify Album/Playlist/Track URL's are supported.", api)
		return "", errors.New("unsupported spotify type of url")
	}
	return savedFile, nil
}
