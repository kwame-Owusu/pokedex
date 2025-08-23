package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type mapResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return fmt.Errorf("Error occurred when making request to location aread endpoint: %w", err)
	}
	defer res.Body.Close()

	var locationArea mapResult
	err = json.NewDecoder(res.Body).Decode(&locationArea)
	if err != nil {
		return fmt.Errorf("Error occured in decoding: %w", err)
	}

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
	os.Exit(0)
	return nil
}
