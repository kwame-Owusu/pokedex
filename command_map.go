package main

import (
	"errors"
	"fmt"
)

type mapResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextURL)
	if err != nil {
		return err
	}

	config.nextURL = locationsResp.Next
	config.prevURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
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
