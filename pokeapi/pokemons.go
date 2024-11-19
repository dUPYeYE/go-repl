package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func (c *Client) ListPokemons(area *string) (RespShallowPokemons, error) {
	if err := godotenv.Load(); err != nil {
		return RespShallowPokemons{}, err
	}

	url := fmt.Sprintf("%s/location-area/%s", os.Getenv("POKEAPI_URL"), *area)

	if val, ok := c.cache.Get(url); ok {
		pokemonsResp := RespShallowPokemons{}
		err := json.Unmarshal(val, &pokemonsResp)
		if err != nil {
			return RespShallowPokemons{}, err
		}

		return pokemonsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	defer res.Body.Close()

	var pokemons RespShallowPokemons
	if err := json.NewDecoder(res.Body).Decode(&pokemons); err != nil {
		return RespShallowPokemons{}, err
	}

	return pokemons, nil
}

func (c *Client) CatchPokemon(pokemon *string) (bool, error) {
	if err := godotenv.Load(); err != nil {
		return false, err
	}

	url := fmt.Sprintf("%s/pokemon/%s", os.Getenv("POKEAPI_URL"), *pokemon)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	var pokemonResp RespPokemon
	if err := json.NewDecoder(res.Body).Decode(&pokemonResp); err != nil {
		return false, err
	}

	caught := rand.Intn(100) > pokemonResp.Experience

	return caught, nil
}
