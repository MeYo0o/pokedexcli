package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/MeYo0o/pokedexcli/internal/api"
	pokecache "github.com/MeYo0o/pokedexcli/internal/cache"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(...string) error
}

type CliConfig struct {
	Next     string
	Previous string
}

// * Commands used in the app
var Commands = map[string]CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	},
	"help": {
		Name:        "help",
		Description: "List all supported commands",
		Callback:    commandHelp,
	},
	"map": {
		Name:        "map",
		Description: "get the location area map ==> next",
		Callback:    commandsMap,
	},
	"mapb": {
		Name:        "mapb",
		Description: "get the location area map <== previous",
		Callback:    commandsMapB,
	},
	"explore": {
		Name:        "explore",
		Description: "get the pokemon inside a location area",
		Callback:    commandExplore,
	},
	"catch": {
		Name:        "catch",
		Description: "catch a pokemon",
		Callback:    commandCatch,
	},
	"inspect": {
		Name:        "inspect",
		Description: "inspect pokedex for a chosen pokemon",
		Callback:    commandInspect,
	},
	"pokedex": {
		Name:        "pokedex",
		Description: "List all Pokemon in your pokedex",
		Callback:    commandPokedex,
	},
}

var commandConfig = CliConfig{
	Next:     "",
	Previous: "",
}

var cache = pokecache.NewCache(5 * time.Second)

var pokedex = make(map[string]api.Pokemon)

// * commands callbacks
func commandExit(params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(params ...string) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandsMap(params ...string) error {
	var url string
	var err error
	var body []byte

	if commandConfig.Next != "" {
		url = commandConfig.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	if _, ok := cache.Get(url); ok {
		body, _ = cache.Get(url)
	} else {
		body, err = apiCall(url)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	var response api.PokeResponse

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return fmt.Errorf("error while decoding the response body:%w", err)
	}

	//* update the config
	commandConfig.Next = response.Next
	commandConfig.Previous = response.Previous

	//* Cache the response
	cache.Add(url, []byte(body))

	for _, locationArea := range response.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandsMapB(params ...string) error {
	var url string
	var err error
	var body []byte

	if commandConfig.Previous != "" {
		url = commandConfig.Previous
	} else {
		return errors.New("you're on the first page")
	}

	if _, ok := cache.Get(url); ok {
		body, _ = cache.Get(url)
	} else {
		body, err = apiCall(url)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	var response api.PokeResponse

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return fmt.Errorf("error while decoding the response body:%w", err)
	}

	//* update the config
	commandConfig.Next = response.Next
	commandConfig.Previous = response.Previous

	//* Cache the response
	cache.Add(url, []byte(body))

	for _, locationArea := range response.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandExplore(params ...string) error {
	var body []byte
	var err error

	if len(params) < 2 {
		return errors.New("please provide a location area to explore")
	}

	firstParamAfterCommand := params[1]

	var url = fmt.Sprint("https://pokeapi.co/api/v2/location-area/", firstParamAfterCommand)

	body, found := cache.Get(url)
	if !found {
		if body, err = apiCall(url); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	var pokemonEncountersResponse api.PokemonEncountersResponse

	err = json.Unmarshal(body, &pokemonEncountersResponse)
	if err != nil {
		return fmt.Errorf("error while decoding the response body:%w", err)
	}

	for _, pokemonEncounter := range pokemonEncountersResponse.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(params ...string) error {
	var body []byte
	var err error

	if len(params) < 2 {
		return errors.New("please provide a pokemon to catch")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", params[1])

	body, err = apiCall(url)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	var pokemonResponse api.Pokemon

	err = json.Unmarshal(body, &pokemonResponse)
	if err != nil {
		return fmt.Errorf("error while decoding the response body:%w", err)
	}

	pokemonName := params[1]

	//* show catching notifier
	fmt.Printf("Throwing a Pokeball at %s...\n", params[1])

	// Seed rand (only once, usually in init())
	rand.Seed(time.Now().UnixNano())

	// Simple formula: chance decreases as BaseExperience increases
	chance := rand.Intn(100)
	difficulty := pokemonResponse.BaseExperience / 10 // tweak divisor to balance

	if chance > difficulty {
		fmt.Printf("%s was caught!\n", pokemonName)
		pokedex[pokemonName] = pokemonResponse
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(params ...string) error {
	if len(params) < 2 {
		return errors.New("please provide a pokemon to inspect")
	}

	pokemonName := params[1]

	if pokemon, ok := pokedex[pokemonName]; ok {
		fmt.Printf(`
Name: %s
Height: %d
Weight: %d
Stats:
  -hp: %d
  -attack: %d
  -defense: %d
  -special-attack: %d
  -special-defense: %d
  -speed: %d
Types:
`,
			pokemon.Name,
			pokemon.Height,
			pokemon.Weight,
			pokemon.PokemonStats[0].BaseStat,
			pokemon.PokemonStats[1].BaseStat,
			pokemon.PokemonStats[2].BaseStat,
			pokemon.PokemonStats[3].BaseStat,
			pokemon.PokemonStats[4].BaseStat,
			pokemon.PokemonStats[5].BaseStat,
		)

		// * Print types separately
		for _, pokemonType := range pokemon.PokemonTypes {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}

func commandPokedex(...string) error {
	if len(pokedex) == 0 {
		return errors.New("your pokedex is empty")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}

// ! Handler
var apiCall = func(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error while getting the location area:%w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while getting the response body:%w", err)
	}
	defer resp.Body.Close()

	return body, nil
}
