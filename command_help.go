package main

import "fmt"

func commandHelp(c *Config, args ...string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for name, cmd := range commands {
        fmt.Printf("%s: %s\n", name, cmd.description)
    }
    return nil
}
