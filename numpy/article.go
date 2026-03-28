package main

import (
	"strings"
)

func fixArticles(s string) string {

	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {

		if strings.ToLower(words[i]) == "a" {

			next := strings.ToLower(words[i+1])

			if strings.HasPrefix(next, "a") ||
				strings.HasPrefix(next, "e") ||
				strings.HasPrefix(next, "i") ||
				strings.HasPrefix(next, "o") ||
				strings.HasPrefix(next, "u") ||
				strings.HasPrefix(next, "h") {

				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}
