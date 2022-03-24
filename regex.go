package main

import "regexp"

func ProcessMessage(sentText string) (string, bool) {
	matched, _ := regexp.MatchString(`(?s)^https?:\/\/open\.spotify\.com\/track/(.*?)(\s*\?si=)`, sentText)
	if !matched {
		return "", false
	}

	re := regexp.MustCompile(`(?s)^https?:\/\/open\.spotify\.com\/track/(.*?)(\s*\?si=)`)

	matches := re.FindAllStringSubmatch(sentText, -1)
	return matches[0][1], true
}
