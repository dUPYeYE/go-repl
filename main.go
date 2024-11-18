package main

import (
	"time"

	"github.com/dUPYeYE/go-repl/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
