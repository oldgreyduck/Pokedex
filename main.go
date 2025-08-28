package main

import (
	  "fmt"
	  "os"
	  "bufio"
	  "strings"
    	  "pokedex/internal/pokeapi"
)

var commands map[string]cliCommand


func main() {
	  scanner:= bufio.NewScanner(os.Stdin)
    commands = map[string]cliCommand{
    "exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
    "help": {name: "help", description: "Displays a help message", callback: commandHelp},
    "map": {name: "map", description: "Displays map locations", callback: commandMap},
    "mapb": {name: "mapb", description: "Displays the previous map locations", callback: commandMapb}}

    config := Config{
        Client:    pokeapi.NewPokeApiClient(),
        Next:      "https://pokeapi.co/api/v2/location-area/",
        Previous:  nil,
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
        if cmd, ok := commands[name]; ok {
            if err := cmd.callback(&config); err != nil {
                fmt.Println(err)
            }
        } else {
            fmt.Println("Unknown command")
        }
	  }
}

type cliCommand struct {
    name        string
    description string
    callback    func(*Config) error
}

type Config struct {
    Client        *pokeapi.PokeApiClient
    Next          string
    Previous      *string
}

