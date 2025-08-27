package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type pokeapiClient struct {
	httpClient http.Client
	cache      map[string][]byte
}

func NewPokeapiClient() *pokeapiClient {
	return &pokeapiClient{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
		cache: make(map[string][]byte),
	}
}

func (c *pokeapiClient) GetLocationAreas(locationURL string) (LocationAreaResp, error) {
	resp, err := c.httpClient.Get(locationURL)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
i		return LocationAreaResp{}, err
	}

	var locationArea LocationAreaResp
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationArea, nil
}
