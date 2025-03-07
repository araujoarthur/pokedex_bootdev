package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}
