package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, args ...string) error {
	pokemon := args[0]
	if _, ok := config.pokedex.Pokemons[pokemon]; !ok {
		err := errors.New("you have not caught that pokemon")
		return err
	}

	inspectionResp, err := config.pokeapiClient.Inspect(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Name : %s\n", pokemon)
	fmt.Printf("Base Experience: %d\n", inspectionResp.BaseExp)
	fmt.Printf("Height: %d\n", inspectionResp.Height)
	fmt.Printf("Weight: %d\n", inspectionResp.Weight)
	fmt.Println("Stats: ")
	for _, stat := range inspectionResp.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range inspectionResp.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(config *config, args ...string) error {
	pokedex := config.pokedex.Pokemons
	if len(pokedex) == 0 {
		return errors.New("No pokemon in pokedex, go catch some pokemons")
	}

	for _, val := range pokedex {
		fmt.Printf("- %s", val.Name)
	}
	return nil
}
