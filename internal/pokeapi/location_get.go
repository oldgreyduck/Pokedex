package pokeapi

import (
        "encoding/json"
        "io"
)

func (c *PokeApiClient) GetLocation(locationName string) (Location, error) {
    locationURL := "https://pokeapi.co/api/v2/location-area/" + locationName
    
    if cachedData, found := c.cache.Get(locationURL); found {
        var location Location
        err := json.Unmarshal(cachedData, &location)
        if err != nil {
            return Location{}, err
        }
        return location, nil
    }

    resp, err := c.httpClient.Get(locationURL)
    if err != nil {
        return Location{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return Location{}, err
    }

    var location Location
    err = json.Unmarshal(body, &location)
    if err != nil {
        return Location{}, err
    }

    c.cache.Add(locationURL, body)

    return location, nil
}
