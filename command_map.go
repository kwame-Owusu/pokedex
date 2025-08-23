package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config) error {
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

func commandMapBack(config *config) error {
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
