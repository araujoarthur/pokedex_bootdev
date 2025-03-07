package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonLower := strings.ToLower(args[0])
	pokemonData, err := c.pokeapiClient.GetPokemon(pokemonLower)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemonData.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonLower)
	if res > 40 {
		fmt.Printf("%s escaped\n", pokemonData.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonData.Name)
	c.caughtPokemon[pokemonData.Name] = pokemonData
	return nil

}
