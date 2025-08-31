package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	minChance = 20
	maxChance = 90
	minExp    = 40  // roughly caterpie
	maxExp    = 306 //rougly mewtwo
)

func (c *Client) Catch(pokemon string, pokedex *Pokedex) (bool, error) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
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

	roll := c.rng.Intn(100)
	if roll < catchChance(pokemonExp.BaseExp) {
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Println("You may now inspect it with the inspect command.")
		if _, ok := pokedex.Pokemons[pokemon]; !ok {
			pokedex.Pokemons[pokemon] = Pokemon{Name: pokemon}
		}
		return true, nil
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
		return false, nil
	}
}

func catchChance(baseExp int) int {
	e := baseExp
	if e < minExp {
		e = minExp
	} else if e > maxExp {
		e = maxExp
	}
	// normalize
	t := float64(e-minExp) / float64(maxExp-minExp)

	// interpolate higher exp -> lower chance
	chanceF := float64(maxChance) - t*float64(maxChance-minChance)

	// to int
	chance := int(chanceF)
	if chance < minChance {
		chance = minChance
	} else if chance > maxChance {
		chance = maxChance
	}

	return chance
}
