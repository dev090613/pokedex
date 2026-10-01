// command_map.go
package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

func commandMap(cfg *config) error {
	var url string
	if cfg.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.Next
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var result struct {
		Next		*string	`json:"next"`
		Previous	*string	`json:"previous"`
		Results		[]struct{
			Name	string	`json:"name"`
			URL		string	`json:"url"`
		} `json:"results"`
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return err
	}

	// fmt.Println(url)

	for _, res := range result.Results {
		fmt.Println(res.Name)
	}

	cfg.Next = result.Next
	cfg.Previous = result.Previous

	return nil
}
