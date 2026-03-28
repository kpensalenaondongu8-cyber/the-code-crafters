package main

import (
	"strings"
	"unicode"
)

func fixPunctuation(s string) string {
	replacer := strings.NewReplacer(
		" ,", ",",
		" .", ".",
		" !", "!",
		" ?", "?",
		" :", ":",
		" ;", ";",
	)
	s = replacer.Replace(s)

	var builder strings.Builder
	runes := []rune(s)
	puncts := ",.!?:;"

	for i := 0; i < len(runes); i++ {
		builder.WriteRune(runes[i])

		if strings.ContainsRune(puncts, runes[i]) {
			if i+1 < len(runes) {
				next := runes[i+1]
				// Add a space if next is a letter, digit, or a quote
				if unicode.IsLetter(next) || unicode.IsDigit(next) || next == '\'' {
					builder.WriteRune(' ')
				}
			}
		}
	}

	return strings.TrimSpace(builder.String())
}

func fixQuotes(s string) string {
	runes := []rune(s)
	var builder strings.Builder
	insideQuotes := false

	for i := 0; i < len(runes); i++ {
		if runes[i] == '\'' {
			if !insideQuotes {
				// Opening quote: trim any space before it
				trimmed := strings.TrimRight(builder.String(), " ")
				builder.Reset()
				builder.WriteString(trimmed)
				builder.WriteRune('\'')
				insideQuotes = true
				// Skip any space immediately after opening quote
				if i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
			} else {
				// Closing quote: trim any space before it
				trimmed := strings.TrimRight(builder.String(), " ")
				builder.Reset()
				builder.WriteString(trimmed)
				builder.WriteRune('\'')
				insideQuotes = false
			}
		} else {
			builder.WriteRune(runes[i])
		}
	}

	return builder.String()
}

// fixColonQuote ensures a space between punctuation and an opening quote
func fixColonQuote(s string) string {
	replacer := strings.NewReplacer(
		":'", ": '",
		",'", ", '",
		".'", ". '",
		"!'", "! '",
		"?'", "? '",
		";'", "; '",
	)
	return replacer.Replace(s)
}
