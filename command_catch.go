package main

import "fmt"

func commandCatch(cfg *config, args []string) error {
	fmt.Println()

	if len(args) == 0 {
		fmt.Println("Please provide a Pokemon to catch")
		return nil
	}

	fmt.Println("Throwing a Pokeball at " + args[0] + "...")
	caught, err := cfg.pokeapiClient.CatchPokemon(&args[0])
	if err != nil {
		fmt.Println(err)
	}

	if caught {
		fmt.Println(args[0] + " was caught!")
	} else {
		fmt.Println(args[0] + " escaped!")
	}

	fmt.Println()
	return nil
}
