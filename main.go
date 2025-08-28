package main

import (
	"github.com/kwame-Owusu/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokedex := pokeapi.Pokedex{
		Pokemons: make(map[string]pokeapi.Pokemon),
	}
	config := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}

	startRepl(config)

}
