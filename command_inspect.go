package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	fmt.Println()
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	if _, ok := cfg.pokedex[args[0]]; !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", cfg.pokedex[args[0]].Name)
	fmt.Printf("Height: %d\n", cfg.pokedex[args[0]].Height)
	fmt.Printf("Weight: %d\n", cfg.pokedex[args[0]].Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range cfg.pokedex[args[0]].Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, t := range cfg.pokedex[args[0]].Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	fmt.Println()
	return nil
}
