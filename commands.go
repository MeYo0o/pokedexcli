package main

var commands = map[string]cliCommands{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "show the supported commands",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "display the next 20 location maps of pokemon world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "display the previous 20 location maps of pokemon world",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "explore the pokemon world",
		callback:    commandExplore,
	},
}
