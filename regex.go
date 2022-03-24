package main

import "regexp"

// ProcessMessage is a function that extracts URI from Spotify URL
func ProcessMessage(sentText string) (string, bool) {
	matched, _ := regexp.MatchString(`(?s)^https?:\/\/open\.spotify\.com\/track/(.*?)(\s*\?si=)`, sentText)
	if !matched {
		return "", false
	}

	re := regexp.MustCompile(`(?s)^https?:\/\/open\.spotify\.com\/track/(.*?)(\s*\?si=)`)

	matches := re.FindAllStringSubmatch(sentText, -1)
	return matches[0][1], true
}
