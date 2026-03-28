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

				if unicode.IsLetter(next) || unicode.IsDigit(next) || next == '\'' {
					builder.WriteRune(' ')
				}
			}
		}
	}

	return strings.TrimSpace(builder.String())
}
