package pokeapi

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Encounters struct {
	AreaName          string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string
}

type PokemonDetails struct {
	BaseExp int `json:"base_experience"`
	Height  int `json:"height"`
	Weight  int `json:"weight"`
}

type Pokedex struct {
	Pokemons map[string]Pokemon
}
