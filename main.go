package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/MeYo0o/pokedexcli/internal/pokecache"
)

func main() {
	// Initialize cache with 5 minute interval
	cache := pokecache.NewCache(5 * time.Minute)

	// config to store the pagination state [previous & next] pages.
	var config config = config{
		Cache: cache,
	}

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

		// Parse the input into command and arguments
		words := cleanInput(userInput)
		if len(words) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}

		commandName := words[0]
		args := words[1:]

		// find if it contains a command == userInput to execute
		if command, ok := commands[commandName]; ok {
			err := command.callback(&config, args...)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
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
