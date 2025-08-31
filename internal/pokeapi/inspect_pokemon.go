package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Inspect(pokemon string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + pokemon

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	pokemonDetails := PokemonDetails{}
	err = json.Unmarshal(data, &pokemonDetails)
	return pokemonDetails, nil
}
