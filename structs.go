package main

import "github.com/MeYo0o/pokedexcli/internal/pokecache"

type cliCommands struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

type pokemonLocationsResponse struct {
	Count int `json:"count"`
	config
	Results []pokemonLocationData `json:"results"`
}

type config struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Cache    *pokecache.Cache
}

type pokemonLocationData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationAreaResponse struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon pokemonData `json:"pokemon"`
}

type pokemonData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
