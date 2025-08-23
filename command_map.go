package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.nextURL != "" {
		url = config.nextURL
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error occurred when making request to location aread endpoint: %w", err)
	}
	defer res.Body.Close()

	var locationArea mapResult
	err = json.NewDecoder(res.Body).Decode(&locationArea)
	if err != nil {
		return fmt.Errorf("Error occured in decoding: %w", err)
	}

	config.nextURL = locationArea.Next
	config.prevURL = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *config) error {
	if config.prevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := config.prevURL

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error occurred when making request to location aread endpoint: %w", err)
	}
	defer res.Body.Close()

	var locationArea mapResult
	err = json.NewDecoder(res.Body).Decode(&locationArea)
	if err != nil {
		return fmt.Errorf("Error occured in decoding: %w", err)
	}

	config.nextURL = locationArea.Next
	config.prevURL = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
	return nil
}
