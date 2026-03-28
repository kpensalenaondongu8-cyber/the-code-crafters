package main

import (
	"regexp"
	"strconv"
	"strings"
)

func ProcessCase(words []string) []string {
	re := regexp.MustCompile(`\((up|low|cap)(, ?(\d+))?\)`)

	for i := 0; i < len(words); i++ {
		match := re.FindStringSubmatch(words[i])
		if len(match) > 0 {
			action := match[1]
			n := 1
			if match[3] != "" {
				n, _ = strconv.Atoi(match[3])
			}

			for j := 0; j < n && i-1-j >= 0; j++ {
				switch action {
				case "up":
					words[i-1-j] = strings.ToUpper(words[i-1-j])
				case "low":
					words[i-1-j] = strings.ToLower(words[i-1-j])
				case "cap":
					words[i-1-j] = strings.Title(words[i-1-j])
				}
			}
			words[i] = ""
		}
	}
	return words
}
