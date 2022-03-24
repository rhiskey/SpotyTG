package spotifydl

import (
	"fmt"
	"github.com/rhiskey/spotytg/utils"
	"github.com/zmb3/spotify/v2"
	"os"
	"os/exec"
)

// Downloader is a function to download files
func Downloader(url string, track spotify.SimpleTrack) string {
	nameTag := fmt.Sprintf("%s.mp3", track.Name)

	ytdlCmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3",
		"-o", track.Name+".%(ext)s", "--audio-quality", "0", url)
	_, err := ytdlCmd.Output()
	if err != nil {
		fmt.Println("=> An error occured while trying to download using youtube-dl")
		fmt.Println("Make sure you have youtube-dl and ffmpeg installed on this system. This was the command we tried to run:")
		fmt.Println(ytdlCmd.String())
		os.Exit(1)
	}

	// Tag the file with metadataa
	return utils.TagFileWithSpotifyMetadata(nameTag, track)
}
