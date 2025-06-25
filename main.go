package main

import (
	"strings"
)

func main() {
	cleanInput("  hello  world  ")
}

// cleanInput splits the input text into words, handling all whitespace characters
func cleanInput(text string) []string {
	return strings.Fields(text)
}
