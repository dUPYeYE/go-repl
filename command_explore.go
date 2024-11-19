package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	fmt.Println()

	if len(args) == 0 {
		fmt.Println("Please provide an area to explore")
		return nil
	}

	pokemons, err := cfg.pokeapiClient.ListPokemons(&args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Exploring " + args[0] + "...")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons.Encounter {
		fmt.Println(pokemon.Pokemon.Name)
	}

	fmt.Println()
	return nil
}
