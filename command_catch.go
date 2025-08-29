package main

import (
	"errors"
	"fmt"
	"math/rand"
	"pokedex/internal/pokeapi"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}
	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	p, err := c.Client.GetPokemon(name)
	if err != nil {
		return err
	}

	// Simple odds based on base experience; tune as you like
	// Lower base exp = easier catch.
	difficulty := p.BaseExperience
	if difficulty < 50 {
		difficulty = 50
	}
	if difficulty > 300 {
		difficulty = 300
	}
	// Threshold controls how many tries on average.
	threshold := 60
	if rand.Intn(difficulty) < threshold {
		// ensure pokedex exists
		if c.Pokedex == nil {
			c.Pokedex = make(map[string]pokeapi.Pokemon)
		}
		c.Pokedex[p.Name] = p
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
