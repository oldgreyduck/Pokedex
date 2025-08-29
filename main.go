package main

import (
	  "fmt"
	  "os"
	  "bufio"
	  "strings"
	  "time"
	  "pokedex/internal/pokeapi"
	  "math/rand"
)

var commands map[string]cliCommand

type cliCommand struct {
    name        string
    description string
    callback    func(*Config, ...string) error
}

type Config struct {
    Client        *pokeapi.PokeApiClient
    Next          string
    Previous      *string
    Pokedex	  map[string]pokeapi.Pokemon
}

func main() {
    rand.Seed(time.Now().UnixNano())
    scanner:= bufio.NewScanner(os.Stdin)
    commands = map[string]cliCommand{
    "exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
    "help": {name: "help", description: "Displays a help message", callback: commandHelp},
    "map": {name: "map", description: "Displays map locations", callback: commandMap},
    "mapb": {name: "mapb", description: "Displays the previous map locations", callback: commandMapb},
    "explore": {name: "explore", description: "Explore a location area", callback: commandExplore},
    "catch": {name: "catch", description: "Attempt to catch a Pokemon", callback: commandCatch},
    }

    config := Config{
        Client:    pokeapi.NewPokeApiClient(30 * time.Second),
        Next:      "https://pokeapi.co/api/v2/location-area/",
        Previous:  nil,
	Pokedex:   make(map[string]pokeapi.Pokemon),
    }

	  for {
		    fmt.Print("Pokedex > ")
		    scanner.Scan()
		    inputLine := scanner.Text()
		    lowerInputLine := strings.ToLower(inputLine)
		    userCommand := strings.Fields(lowerInputLine)
        if len(userCommand) == 0 {
            fmt.Println("Unknown command")
            continue
        }
        name := userCommand[0]
        args := userCommand [1:]
        if cmd, ok := commands[name]; ok {
            if err := cmd.callback(&config, args...); err != nil {
                fmt.Println(err)
            }
        } else {
            fmt.Println("Unknown command")
        }
	  }
}

