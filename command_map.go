package main

import "fmt"

func commandMapf(cfg *config, args []string) error {
	fmt.Println()

	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Locations:")
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	fmt.Println()
	return nil
}

func commandMapb(cfg *config, args []string) error {
	fmt.Println()

	if cfg.prevLocationsURL == nil {
		fmt.Println("No previous locations, this is the first page")
		fmt.Println()
		return nil
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Locations:")
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	fmt.Println()
	return nil
}
