package main

import "strings"

func fixQuotes(s string) string {

	runes := []rune(s)

	var builder strings.Builder

	insideQuotes := false

	for i := 0; i < len(runes); i++ {

		if runes[i] == '\'' {

			if !insideQuotes {

				trimmed := strings.TrimRight(builder.String(), " ")
				builder.Reset()
				builder.WriteString(trimmed)

				builder.WriteRune('\'')

				insideQuotes = true

				if i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}

			} else {

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
