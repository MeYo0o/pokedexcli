package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// config to store the pagination state [previous & next] pages.
	var config config = config{}

	// wait for user command
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string

	// intro
	fmt.Println("Welcome to the Pokedex!")

	// prompt the first command prefix
	fmt.Print("Pokedex > ")

	// loop through each userInput after entering "return key"
	for scanner.Scan() {

		// get user input
		userInput = scanner.Text()

		// find if it contains a command == userInput to execute
		if command, ok := commands[userInput]; ok {
			command.callback(&config)
		} else {
			fmt.Println("Unknown command")
		}

		// clear userInput
		userInput = ""

		// prompt the next command prefix
		fmt.Print("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Invalid Input: %w", err)
	}

}
