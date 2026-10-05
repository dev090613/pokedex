package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, val := range cfg.commands {
		fmt.Printf("%s: %s", key, val.description)
		fmt.Println()
	}

	return nil
}

