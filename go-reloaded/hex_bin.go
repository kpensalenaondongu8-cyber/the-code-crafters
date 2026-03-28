package main

import "strconv"

func ProcessHexBin(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			dec, err := strconv.ParseInt(words[i-1], 16, 0)
			if err == nil {
				words[i-1] = strconv.FormatInt(dec, 10)
			}
			words[i] = ""
		} else if words[i] == "(bin)" && i > 0 {
			dec, err := strconv.ParseInt(words[i-1], 2, 0)
			if err == nil {
				words[i-1] = strconv.FormatInt(dec, 10)
			}
			words[i] = ""
		}
	}
	return words
}
