package main

import "strings"

func ProcessArticles(words []string) []string {

	vowels := "aeiouh"

	for i := 0; i < len(words)-1; i++ {
		current := strings.ToLower(words[i])
		next := strings.ToLower(words[i+1])

		if len(next) == 0 {
			continue
		}

		first := string(next[0])

		if current == "a" && strings.ContainsAny(first, vowels) {

			words[i] = "an"
		}

		if current == "an" && !strings.ContainsAny(first, vowels) {
			words[i] = "a"
		}
	}

	return words
}
