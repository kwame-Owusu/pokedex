package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Catch(pokemon string, pokedex *Pokedex) (bool, error) {
	fmt.Printf("Throwing ball at %s...\n", pokemon)
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
	pokemonStruct := Pokemon{Name: pokemon}
	pokedex.Pokemons[pokemon] = pokemonStruct
	for _, val := range pokedex.Pokemons {
		fmt.Printf("- %s\n", val.Name)
	}
	return false, nil
}
