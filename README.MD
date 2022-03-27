# SpotyTG

[![Go](https://github.com/rhiskey/SpotyTG/actions/workflows/go.yml/badge.svg)](https://github.com/rhiskey/SpotyTG/actions/workflows/go.yml)
[![CodeQL](https://github.com/rhiskey/SpotyTG/actions/workflows/codeql.yml/badge.svg)](https://github.com/rhiskey/SpotyTG/actions/workflows/codeql.yml)

Telegram Bot for music import from Spotify

[![giphy.gif](https://media.giphy.com/media/PdK4aHlXmBmOvpPU04/giphy.gif)](https://media.giphy.com/media/PdK4aHlXmBmOvpPU04/giphy.gif)

---

## Installation

### Executable from [Releases](https://github.com/rhiskey/SpotyTG/releases)
#### Prerequisites

* [youtube-dl](http://ytdl-org.github.io/youtube-dl/download.html)
* [ffmpeg](https://ffmpeg.org/download.html)


### From Source (IN PROGRESS)
#### Prerequisites

* [youtube-dl](http://ytdl-org.github.io/youtube-dl/download.html)
* [ffmpeg](https://ffmpeg.org/download.html)

Make sure you have **golang**, **youtube-dl** and **ffmpeg** installed.

`go install github.com/rhiskey/spotytg`


---
### Environment variables to pass:
```bash
DEBUG=false 
ROLLBAR_TOKEN=xxxxxxxxxxxxxxxxxxxx111111
TELEGRAM_APITOKEN=11111111111:XXXXXXXXXXXxxxxxxxxxx
SPOTIFY_ID=xxxxxxxxxxxxxxxxxxxx111111
SPOTIFY_SECRET=xxxxxxxxxxxxxxxxxxxx111111
```

Alternative build executable
```shell
env GOOS=linux GOARCH=amd64 go build .
## OR 
$env:GOOS = "linux" 
go build .
```