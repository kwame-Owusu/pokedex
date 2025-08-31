package main

import "fmt"

func commandInspect(config *config, args ...string) error {
	pokemon := args[0]
	inspectionResp, err := config.pokeapiClient.Inspect(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Name : %s\n", pokemon)
	fmt.Printf("Base Experience: %d\n", inspectionResp.BaseExp)
	fmt.Printf("Height: %d\n", inspectionResp.Height)
	fmt.Printf("Weight: %d\n", inspectionResp.Weight)
	fmt.Println("Stats: ")
	for _, stat := range inspectionResp.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range inspectionResp.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}
