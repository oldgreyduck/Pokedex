package main

import (
	"fmt"
	"sort"
)

func commandPokedex(c *Config, args ...string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Time to catch a Pokemon!")
		return nil
	}
	names := make([]string, 0, len(c.Pokedex))
	for name := range c.Pokedex {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Your Pokedex:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
