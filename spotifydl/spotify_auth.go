package spotifydl

import (
	"github.com/zmb3/spotify/v2"
)

var (
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

// UserData is a struct to hold all variables
type UserData struct {
	UserClient      *spotify.Client
	TrackList       []spotify.FullTrack
	SimpleTrackList []spotify.SimpleTrack
	YoutubeIDList   []string
}
