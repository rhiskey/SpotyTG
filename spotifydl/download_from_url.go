package spotifydl

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"os"
	"strings"
)

// DonwloadFromURL is a function to decide which type of content URL contains and download it
func DonwloadFromURL(spotifyURL string, spotyclient *spotify.Client, ctx context.Context) string {
	var trackID string
	var playlistID string
	var albumID string

	var savedFile string

	if len(spotifyURL) == 0 {
		fmt.Println("=> Spotify URL required.")
		return ""
	}

	splitURL := strings.Split(spotifyURL, "/")

	if len(splitURL) < 2 {
		fmt.Println("=> Please enter the url copied from the spotify client.")
		os.Exit(1)
	}

	spotifyID := splitURL[len(splitURL)-1]
	if strings.Contains(spotifyID, "?") {
		spotifyID = strings.Split(spotifyID, "?")[0]
	}

	if strings.Contains(spotifyURL, "album") {
		albumID = spotifyID
		DownloadAlbum(albumID, spotyclient, ctx)
	} else if strings.Contains(spotifyURL, "playlist") {
		playlistID = spotifyID
		DownloadPlaylist(playlistID, spotyclient, ctx)
	} else if strings.Contains(spotifyURL, "track") {
		trackID = spotifyID
		savedFile = DownloadSong(trackID, spotyclient, ctx)
	} else {
		fmt.Println("=> Only Spotify Album/Playlist/Track URL's are supported.")
	}
	return savedFile
}
