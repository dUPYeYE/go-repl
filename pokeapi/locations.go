package pokeapi

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	if err := godotenv.Load(); err != nil {
		return RespShallowLocations{}, err
	}

	url := os.Getenv("POKEAPI_URL") + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	var locations RespShallowLocations
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return RespShallowLocations{}, err
	}

	return locations, nil
}
