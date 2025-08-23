package main

import (
	"github.com/kwame-Owusu/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)

}
