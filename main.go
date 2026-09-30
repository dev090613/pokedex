package main

import (
    "fmt"
)

type cliCommand struct {
    name		string
    description	string
    callback	func(*config) error
}

type config struct {
    commands	map[string]cliCommand
    Next		*string	`json:"next"`
    Previous	*string	`json:"previous"`
    Results		[]struct{
	Name	string	`json:"name"`
	URL		string	`json:"url"`
    } `json:"results"`
}

func main() {
    endpoint := "location-area"
    url := fmt.Sprintf("https://pokeapi.co/api/v2/%v/", endpoint)

    cfg := &config{
	commands: getCommand(),
    }
    startRepl(cfg)
}
