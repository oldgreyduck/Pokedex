package main

import (
	  "fmt"
	  "os"
	  "bufio"
	  "strings"
)


var commands map[string]cliCommand

func main() {
	  scanner:= bufio.NewScanner(os.Stdin)
    commands = map[string]cliCommand{
    "exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
    "help": {name: "help", description: "Displays a help message", callback: commandHelp},
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
            if err := cmd.callback(); err != nil {
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
    callback    func() error
}

