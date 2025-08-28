package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Catch(pokemon string, pokedex map[string]Pokemon) (bool, error) {
	fmt.Printf("Throwing ball at %s...", pokemon)
	url := baseURL + "/pokemon/" + pokemon

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	pokemonExp := PokemonDetails{}
	err = json.Unmarshal(data, &pokemonExp)
	if err != nil {
		return false, err
	}

	return false, nil
}
