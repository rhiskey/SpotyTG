package spotifydl

import (
	"fmt"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rhiskey/spotytg/utils"
	"github.com/rollbar/rollbar-go"
	"github.com/zmb3/spotify/v2"
	"log"
	"os/exec"
)

// Downloader is a function to download files
func Downloader(url string, track spotify.FullTrack, api *structures.Api) string {
	nameTag := fmt.Sprintf("%s.mp3", track.Name)

	ytdlCmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3",
		"-o", track.Name+".%(ext)s", "--audio-quality", "0", url)
	_, err := ytdlCmd.Output()
	if err != nil {
		utils.LogWithBot("⛔ => An error occured while trying to download using youtube-dl", api)
		utils.LogWithBot("⛔ Make sure you have youtube-dl and ffmpeg installed on this system. This was the command we tried to run:", api)
		utils.LogWithBot(fmt.Sprintf(ytdlCmd.String()), api)
		rollbar.Critical(err)
		log.Fatal(err)
	}

	// Tag the file with metadataa
	//return utils.TagFileWithSpotifyMetadataV2(nameTag, track.SimpleTrack, api)
	return utils.TagFileWithSpotifyMetadata(nameTag, track)
}
