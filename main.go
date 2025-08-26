package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


func main() {
	scanner:= bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputLine := scanner.Text()
		lowerInputLine := strings.ToLower(inputLine)
		command := strings.Fields(lowerInputLine)
		if len(command) > 0 {
			fmt.Printf("Your command was: %s\n", command[0])
		} else {
			fmt.Printf("Your command was: (empty input)\n")
		}
	}
}
