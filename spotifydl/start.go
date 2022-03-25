package spotifydl

import (
	"context"
	"errors"
	"fmt"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rhiskey/spotytg/utils"
	"github.com/zmb3/spotify/v2"
	"log"
	"strings"
)

// DownloadPlaylist Start initializes complete program
func DownloadPlaylist(pid string, api *structures.Api, ctx context.Context) {
	user := api.SpotifyClient
	cli := structures.UserData{
		UserClient: user,
	}
	playlistID := spotify.ID(pid)

	trackListJSON, err := cli.UserClient.GetPlaylistTracks(ctx, playlistID)
	if err != nil {
		fmt.Println("⚠ Playlist not found!")
		//os.Exit(1)
	}
	for _, val := range trackListJSON.Tracks {
		cli.TrackList = append(cli.TrackList, val.Track)
	}

	for page := 0; ; page++ {
		err := cli.UserClient.NextPage(ctx, trackListJSON)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, val := range trackListJSON.Tracks {
			cli.TrackList = append(cli.TrackList, val.Track)
		}
	}

	DownloadTrackList(cli, api)
}

// DownloadAlbum Download album according to
func DownloadAlbum(aid string, api *structures.Api, ctx context.Context) {
	user := api.SpotifyClient
	cli := structures.UserData{
		UserClient: user,
	}
	albumID := spotify.ID(aid)
	album, err := user.GetAlbum(ctx, albumID)
	if err != nil {
		utils.LogWithBot("⚠ Album not found!", api)
	}
	for _, val := range album.Tracks.Tracks {
		cli.TrackList = append(cli.TrackList, spotify.FullTrack{
			SimpleTrack: val,
			Album:       album.SimpleAlbum,
		})
	}
	DownloadTrackList(cli, api)
}

// DownloadSong will download a song with its identifier
func DownloadSong(sid string, api *structures.Api, ctx context.Context) (string, error) {
	user := api.SpotifyClient
	cli := structures.UserData{
		UserClient: user,
	}
	songID := spotify.ID(sid)
	song, err := user.GetTrack(ctx, songID)
	if err != nil {
		utils.LogWithBot("⚠ Song not found!", api)
		return "", errors.New("song not found")
	}

	cli.TrackList = append(cli.TrackList, spotify.FullTrack{
		SimpleTrack: song.SimpleTrack,
		Album:       song.Album,
	})
	return DownloadTrackList(cli, api), nil
}

// DownloadTrackList Start downloading given list of tracks
func DownloadTrackList(cli structures.UserData, api *structures.Api) string {
	var savedFile string
	//utils.LogWithBot(fmt.Sprintf("🎵 Found ", len(cli.TrackList), " tracks"), api)
	//utils.LogWithBot("🔎 Searching and downloading tracks", api)
	for _, val := range cli.TrackList {
		var artistNames []string
		for _, artistInfo := range val.Artists {
			artistNames = append(artistNames, artistInfo.Name)
		}
		searchTerm := strings.Join(artistNames, " ") + " " + val.Name
		youtubeID, err := GetYoutubeId(searchTerm, val.Duration/1000)
		if err != nil {
			utils.LogWithBot(fmt.Sprintf("⚠ Error occured for %s error: %s", val.Name, err), api)
			continue
		}
		cli.YoutubeIDList = append(cli.YoutubeIDList, youtubeID)
	}
	for index, track := range cli.YoutubeIDList {
		fmt.Println()
		ytURL := "https://www.youtube.com/watch?v=" + track
		utils.LogWithBot(fmt.Sprintf("🔄️ Downloading: "+cli.TrackList[index].Name), api)
		savedFile = Downloader(ytURL, cli.TrackList[index].SimpleTrack, api)
		fmt.Println()
	}
	//utils.LogWithBot("✔ Download complete!", api)

	return savedFile
}
