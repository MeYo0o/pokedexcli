package main

import (
	"strings"
)

func main() {
	cleanInput("  hello  world  ")
}

func cleanInput(text string) []string {
	stringsSli := strings.Split(text, " ")
	var splitString []string
	for _, word := range stringsSli {
		if word == " " || word == "" {
			continue
		}

		splitString = append(splitString, word)
	}

	return splitString

}
