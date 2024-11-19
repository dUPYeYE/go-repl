package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	fmt.Println()

	if len(cfg.pokedex) == 0 {
		return errors.New("You have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf(" - %s\n", name)
	}

	fmt.Println()
	return nil
}
