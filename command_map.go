package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config, args ...string) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextURL)
	if err != nil {
		return err
	}

	config.nextURL = locationsResp.Next
	config.prevURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *config, args ...string) error {
	if config.prevURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.prevURL)
	if err != nil {
		return err
	}

	config.nextURL = locationResp.Next
	config.prevURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(config *config, args ...string) error {
	if len(args) > 1 || len(args) == 0 {
		return errors.New("Please check the length of the args, can only explore one location")
	}

	location := args[0]
	encounters, err := config.pokeapiClient.Explore(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)

	}
	return nil
}
