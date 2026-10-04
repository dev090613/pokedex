package main

import "github.com/dev090613/pokedexcli/internal/pokecache"

type cliCommand struct {
    name		string
    description	string
    callback	func(*config) error
}

type config struct {
    commands	map[string]cliCommand
    Next		*string	`json:"next"`
    Previous	*string	`json:"previous"`
	Cache 		*pokecache.Cache
}

