package api

type PokeResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncountersResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonName `json:"pokemon"`
}

type PokemonName struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	PokemonStats   []PokemonStat `json:"stats"`
	PokemonTypes   []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}
