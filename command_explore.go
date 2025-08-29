package main

import (
        "fmt"
        "errors"
)

func commandExplore(c *Config, args ...string) error {
        if len(args) == 0 {
                return errors.New("No location entered")
        }
        locationName := args[0]
        res, err := c.Client.GetLocation(locationName)
        if err != nil {
                return err
        }
        fmt.Printf("Exploring %s...\n", locationName)
        fmt.Println("Found Pokemon:")
        for _, encounter := range res.PokemonEncounters {
                fmt.Printf(" - %s\n", encounter.Pokemon.Name)
        }
        return nil
}
