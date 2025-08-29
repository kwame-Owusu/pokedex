package main

import (
	"github.com/kwame-Owusu/pokedex/internal/pokeapi"
	"math/rand"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	pokeClient := pokeapi.NewClient(5*time.Second, rng)
	pokedex := pokeapi.Pokedex{
		Pokemons: make(map[string]pokeapi.Pokemon),
	}
	config := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}

	startRepl(config)

}
