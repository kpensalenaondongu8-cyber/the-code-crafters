package main

import "strings"

func Tokenize(text string) []string {
	var tokens []string
	current := ""

	for _, r := range text {
		switch {
		case r == ' ' || r == '\n' || r == '\t':
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
		case strings.ContainsRune(".,!?;:()\"", r):
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(r))
		default:
			current += string(r)
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	// Combine parentheses groups like (this is a group)
	var combinedTokens []string
	i := 0
	for i < len(tokens) {
		if tokens[i] == "(" {
			combined := tokens[i]
			i++
			for i < len(tokens) && tokens[i] != ")" {
				combined += tokens[i]
				i++
			}
			if i < len(tokens) {
				combined += tokens[i]
			}
			combinedTokens = append(combinedTokens, combined)
			i++
		} else {
			combinedTokens = append(combinedTokens, tokens[i])
			i++
		}
	}

	return combinedTokens
}
