package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MeYo0o/pokedexcli/internal/commands"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userCommand := strings.ToLower(scanner.Text())

		params := getParams(userCommand)

		if command, ok := commands.Commands[params[0]]; ok {
			err := command.Callback(params...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func getParams(input string) []string {
	params := strings.Split(strings.TrimSpace(input), " ")

	return params
}
