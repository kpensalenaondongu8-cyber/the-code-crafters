package main

import (
	"strconv"
	"strings"
)

func ProcessText(text string) string {

	text = strings.ReplaceAll(text, "(up, ", "(up,")
	text = strings.ReplaceAll(text, "(low, ", "(low,")
	text = strings.ReplaceAll(text, "(cap, ", "(cap,")

	words := strings.Fields(text)

	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(hex)" {
			if len(result) > 0 {
				val, err := strconv.ParseInt(result[len(result)-1], 16, 64)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(val, 10)
				}
			}

		} else if word == "(bin)" {
			if len(result) > 0 {
				val, err := strconv.ParseInt(result[len(result)-1], 2, 64)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(val, 10)
				}
			}

		} else if strings.HasPrefix(word, "(up") {
			result = applyCase(result, word, strings.ToUpper)

		} else if strings.HasPrefix(word, "(low") {
			result = applyCase(result, word, strings.ToLower)

		} else if strings.HasPrefix(word, "(cap") {
			result = applyCase(result, word, capitalize)

		} else {
			result = append(result, word)
		}
	}

	final := strings.Join(result, " ")
	final = fixPunctuation(final)
	final = fixArticles(final)
	final = fixQuotes(final)
	final = fixColonQuote(final)

	return final
}
