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

go install github.com/rhiskey/spotytg
---
### Environment variables to pass:
```bash
DEBUG=false 
ROLLBAR_TOKEN=4a505f108ef0402b9b2db8f363793946
TELEGRAM_APITOKEN=11111111111:XXXXXXXXXXXxxxxxxxxxx
SPOTIFY_ID=05c5bdfd503c4a489961d2e63f917b1d
SPOTIFY_SECRET=xxxxxxxxxxxxxxxxxxxx111111
```