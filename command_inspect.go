// go
package main

import (
	"errors"
	"fmt"
	"sort"
)

func commandInspect(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon>")
	}
	key := args[0]

	p, ok := c.Pokedex[key]
	if !ok {
		fmt.Printf("You haven't caught %s yet\n", key)
		return nil
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Printf("Base Experience: %d\n", p.BaseExperience)

	// Types (sorted by slot)
	sort.Slice(p.Types, func(i, j int) bool { return p.Types[i].Slot < p.Types[j].Slot })
	fmt.Printf("Types:\n")
	for _, t := range p.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	// Stats
	fmt.Printf("Stats:\n")
	for _, s := range p.Stats {
		fmt.Printf(" - %s: %d\n", s.Stat.Name, s.BaseStat)
	}

	// Abilities
	fmt.Printf("Abilities:\n")
	for _, a := range p.Abilities {
		hidden := ""
		if a.IsHidden {
			hidden = " (hidden)"
		}
		fmt.Printf(" - %s%s\n", a.Ability.Name, hidden)
	}

	return nil
}
