package main

import "strings"

func cleanInput(text string) []string {
    lowercased := strings.ToLower(text)
    words := strings.Fields(lowercased)
    return words
}

