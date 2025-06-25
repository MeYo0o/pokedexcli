package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// wait for user command
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string

	// intro
	fmt.Print("Pokedex >")

	// loop through each userInput after entering "return key"
	for scanner.Scan() {

		// get user input
		userInput = scanner.Text()

		// get each word separated
		textSli := cleanInput(userInput)

		// clear userInput
		userInput = ""

		if len(textSli) != 0 {
			// print just the first word
			fmt.Printf("Your command was: %s\n", strings.ToLower(textSli[0]))
		} else {
			fmt.Println("Invalid Input")
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Invalid Input: %w", err)
	}

}

// cleanInput splits the input text into words, handling all whitespace characters
func cleanInput(text string) []string {
	return strings.Fields(text)
}
