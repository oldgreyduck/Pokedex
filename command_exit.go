package main

import (
  "os"
  "fmt"
)

func commandExit(c *Config) error {
  fmt.Println("Closing the Pokedex... Goodbye!")
  os.Exit(0)
  return nil
}
