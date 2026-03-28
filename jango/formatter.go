package main

import (
	"strconv"
	"strings"
	"unicode"
)

// ProcessText applies all transformations to the input text
func ProcessText(text string) string {

	text = strings.ReplaceAll(text, "(up, ", "(up,")
	text = strings.ReplaceAll(text, "(low, ", "(low,")
	text = strings.ReplaceAll(text, "(cap, ", "(cap,")

	words := strings.Fields(text) // Splits the input text into tokens (individual words/strings)

	var result []string // holds the progressively transformed words

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(hex)" { // Converts previous word from hex
			if len(result) > 0 {
				val, err := strconv.ParseInt(result[len(result)-1], 16, 64)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(val, 10)
				}
			}
		} else if word == "(bin)" { // Converts previous word from bin
			if len(result) > 0 {
				val, err := strconv.ParseInt(result[len(result)-1], 2, 64)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(val, 10)
				}
			}
		} else if strings.HasPrefix(word, "(up") {
			result = applyCase(result, word, strings.ToUpper)
			// IMPORTANT: do NOT append the marker itself
		} else if strings.HasPrefix(word, "(low") {
			result = applyCase(result, word, strings.ToLower)
		} else if strings.HasPrefix(word, "(cap") {
			result = applyCase(result, word, capitalize)
		} else {
			result = append(result, word)
		}
	}

	final := strings.Join(result, " ")
	final = fixPunctuation(final) // Fix punctuation spacing
	final = fixArticles(final)    // Fix 'a' vs 'an'
	final = fixQuotes(final)
	final = fixColonQuote(final)
	return final
}

// applyCase applies case transformation to N previous words
func applyCase(result []string, marker string, fn func(string) string) []string {
	n := 1
	if strings.Contains(marker, ",") {
		parts := strings.Split(marker, ",")
		if len(parts) == 2 {
			num, err := strconv.Atoi(strings.TrimSpace(strings.Trim(parts[1], ")")))
			if err == nil {
				n = num
			}
		}
	}
	for j := len(result) - n; j < len(result); j++ {
		if j >= 0 {
			result[j] = fn(result[j])
		}
	}
	return result
}

// capitalize first letter
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// fixPunctuation ensures punctuation marks are attached to the previous word
// and separated from the next word by a single space (including before quotes).
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

// fixArticles changes "a" to "an" before vowels/h and preserves capitalization.
func fixArticles(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		w := words[i]
		next := strings.ToLower(words[i+1])

		if strings.EqualFold(w, "a") {
			if strings.HasPrefix(next, "a") ||
				strings.HasPrefix(next, "e") ||
				strings.HasPrefix(next, "i") ||
				strings.HasPrefix(next, "o") ||
				strings.HasPrefix(next, "u") ||
				strings.HasPrefix(next, "h") {
				// Preserve capitalization
				if w == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	return strings.Join(words, " ")
}

// fixQuotes ensures that single quotes hug the text inside them without spaces.
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
