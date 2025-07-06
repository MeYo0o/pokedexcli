package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// cleanInput splits the input text into words, handling all whitespace characters
func cleanInput(text string) []string {
	return strings.Fields(text)
}

func commandExit(config *config) error {
	// notify the user
	fmt.Println("Closing the Pokedex... Goodbye!")

	// exit the cli app
	os.Exit(0)

	return nil
}
func commandHelp(config *config) error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)

	return nil
}

func commandMap(config *config) error {
	var url string

	// if we loaded this api already, config != null
	if config != nil && config.Next != "" {
		url = config.Next
	} else {
		// default api endpoint to get the first page [20 locations]
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	// Check cache first
	if cachedData, found := config.Cache.Get(url); found {
		fmt.Println("Using cached data...")

		var pokemonLocationsResponse pokemonLocationsResponse
		if err := json.Unmarshal(cachedData, &pokemonLocationsResponse); err != nil {
			return fmt.Errorf("failed to unmarshal cached response: %w", err)
		}

		// update the config
		*config = pokemonLocationsResponse.config

		// print the results map
		for _, location := range pokemonLocationsResponse.Results {
			fmt.Println(location.Name)
		}

		return nil
	}

	fmt.Println("Fetching data from API...")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create a request: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to receive a response: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	// Cache the response
	config.Cache.Add(url, body)

	var pokemonLocationsResponse pokemonLocationsResponse

	if err := json.Unmarshal([]byte(body), &pokemonLocationsResponse); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// update the config
	*config = pokemonLocationsResponse.config

	// print the results map
	for _, location := range pokemonLocationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	var url string

	if config != nil && config.Previous != "" {
		url = config.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	// Check cache first
	if cachedData, found := config.Cache.Get(url); found {
		fmt.Println("Using cached data...")

		var pokemonLocationsResponse pokemonLocationsResponse
		if err := json.Unmarshal(cachedData, &pokemonLocationsResponse); err != nil {
			return fmt.Errorf("failed to unmarshal cached response: %w", err)
		}

		// update the config
		*config = pokemonLocationsResponse.config

		// print the results map
		for _, location := range pokemonLocationsResponse.Results {
			fmt.Println(location.Name)
		}

		return nil
	}

	fmt.Println("Fetching data from API...")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create a request: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to receive a response: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	// Cache the response
	config.Cache.Add(url, body)

	var pokemonLocationsResponse pokemonLocationsResponse

	if err := json.Unmarshal([]byte(body), &pokemonLocationsResponse); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// update the config
	*config = pokemonLocationsResponse.config

	// print the results map
	for _, location := range pokemonLocationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
