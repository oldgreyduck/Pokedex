package main

import "fmt"

func commandMapb(c *Config) error {
    if c.Previous == nil {
            fmt.Println("you're on the first page")
            return nil
    }
    res, err := c.Client.GetLocationAreas(c.Previous)
        if err != nil {
            return err
        }
        for _, location := range res.Results {
            fmt.Println(location.Name)
        }
        c.Next = res.Next
        c.Previous = res.Previous
        return nil
}
