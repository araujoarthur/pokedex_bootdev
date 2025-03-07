package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	if len(c.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon yet.")
		return nil
	}

	for k := range c.caughtPokemon {
		fmt.Println("-", k)
	}

	return nil
}
