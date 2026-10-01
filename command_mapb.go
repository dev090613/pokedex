package main

import (
	"net/http"
	"encoding/json"
)

func commandMapb(cfg *config) error {

	var url string
	if (cfg.Previous == nil) {
		url = "https://pokeapi.co/api/v2/location-area"
	} else { 
		url = *cfg.Previous
	}

	http.Get(url)

	return nil
}
