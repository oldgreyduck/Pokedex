package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *PokeApiClient) GetLocationAreas(locationURL string) (LocationAreaResp, error) {
  if cachedData, found := c.cache.Get(locationURL); found {
    var locationArea LocationAreaResp
    err := json.Unmarshal(cachedData, &locationArea)
    if err != nil {
        return LocationAreaResp{}, err
    }
    return locationArea, nil
}

  resp, err := c.httpClient.Get(locationURL)
  if err != nil {
    return LocationAreaResp{}, err
  }
  defer resp.Body.Close()

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return LocationAreaResp{}, err
  }

  var locationArea LocationAreaResp
  err = json.Unmarshal(body, &locationArea)
  if err != nil {
    return LocationAreaResp{}, err
  }

c.cache.Add(locationURL, body)

  return locationArea, nil
}

