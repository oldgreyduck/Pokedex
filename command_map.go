package main

import "fmt"

func commandMap(c *Config, args ...string) error {
    res, err := c.Client.GetLocationAreas(c.Next)
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
