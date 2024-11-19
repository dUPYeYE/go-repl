package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	var pokemons RespShallowPokemons
	if err := json.Unmarshal(dat, &pokemons); err != nil {
		return RespShallowPokemons{}, err
	}

	c.cache.Add(url, dat)
	return pokemons, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	if err := godotenv.Load(); err != nil {
		return Pokemon{}, err
	}

	url := fmt.Sprintf("%s/pokemon/%s", os.Getenv("POKEAPI_URL"), pokemonName)

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
