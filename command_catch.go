package main

func commandCatch(config *config, args ...string) error {
	_, err := config.pokeapiClient.Catch(args[0], &config.pokedex)
	if err != nil {
		return err
	}
	return nil
}
