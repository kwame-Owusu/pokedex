package main

import "fmt"

func commandCatch(config *config, args ...string) error {
	catchResp, err := config.pokeapiClient.Catch(args[0], config.pokedex)
	if err != nil {
		return nil
	}
	fmt.Print(catchResp)
	return nil
}
