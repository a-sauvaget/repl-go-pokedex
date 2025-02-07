package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	fmt.Println("Hello, World!")
}