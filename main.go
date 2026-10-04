// main.go
package main

import (
	"time"
	"github.com/dev090613/pokedexcli/internal/pokecache"
)

func main() {
	// endpoint := "location-area"
	// url := fmt.Sprintf("https://pokeapi.co/api/v2/%v/", endpoint)
	c := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		commands: getCommand(),
		Next: nil,
		Previous: nil,
		Cache: c,
	}
	startRepl(cfg)
}
