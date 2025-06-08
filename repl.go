package main

import (
	"strings"
)

func cleanInput(input string) []string {
	words := strings.Fields(input)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return words
}
