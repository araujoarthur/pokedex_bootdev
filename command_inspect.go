package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}

	plower := strings.ToLower(args[0])

	pokemon, exists := c.caughtPokemon[plower]
	if !exists {
		return errors.New("you have not caught this pokemon yet")
	}

	template := `Name: %s
Height: %d
Weight: %d
Stats:
%s
Types:
%s
	`
	statsString := ""
	for _, stat := range pokemon.Stats {
		statsString = statsString + fmt.Sprintf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	typesString := ""
	for _, typ := range pokemon.Types {
		typesString = typesString + fmt.Sprintf("\t-%s\n", typ.Type.Name)
	}

	builtString := fmt.Sprintf(template, pokemon.Name, pokemon.Height, pokemon.Weight, statsString, typesString)

	fmt.Println(builtString)

	return nil
}
