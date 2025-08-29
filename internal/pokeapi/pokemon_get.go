package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *PokeApiClient) GetPokemon(name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	if cached, ok := c.cache.Get(url); ok {
		var p Pokemon
		if err := json.Unmarshal(cached, &p); err != nil {
			return Pokemon{}, err
		}
		return p, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Pokemon{}, fmt.Errorf("pokemon not found: %s", name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var p Pokemon
	if err := json.Unmarshal(body, &p); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return p, nil
}
