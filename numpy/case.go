package main

import (
	"strconv"
	"strings"
)

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

func capitalize(s string) string {

	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
