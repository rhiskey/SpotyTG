package spotifydl

import (
	"fmt"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rhiskey/spotytg/utils"
	"github.com/rollbar/rollbar-go"
	"github.com/zmb3/spotify/v2"
	"os"
	"os/exec"
)

// Downloader is a function to download files
func Downloader(url string, track spotify.SimpleTrack, api *structures.Api) string {
	nameTag := fmt.Sprintf("%s.mp3", track.Name)

	ytdlCmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3",
		"-o", track.Name+".%(ext)s", "--audio-quality", "0", url)
	_, err := ytdlCmd.Output()
	if err != nil {
		utils.LogWithBot("⛔ => An error occured while trying to download using youtube-dl", api)
		//fmt.Println("=> An error occured while trying to download using youtube-dl")
		utils.LogWithBot("⛔ Make sure you have youtube-dl and ffmpeg installed on this system. This was the command we tried to run:", api)
		//fmt.Println("Make sure you have youtube-dl and ffmpeg installed on this system. This was the command we tried to run:")
		//fmt.Println(ytdlCmd.String())
		utils.LogWithBot(fmt.Sprintf(ytdlCmd.String()), api)
		rollbar.Critical(err)
		os.Exit(1)
	}

	// Tag the file with metadataa
	return utils.TagFileWithSpotifyMetadata(nameTag, track, api)
}
