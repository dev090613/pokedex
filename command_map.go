// command_map.go
package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

func commandMap(cfg *config, args ...string) error {
	var url string
	if cfg.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.Next
	}

	var result struct {
		Next		*string	`json:"next"`
		Previous	*string	`json:"previous"`
		Results		[]struct{
			Name	string	`json:"name"`
			URL		string	`json:"url"`
		} `json:"results"`
	}

	raw, hit := cfg.Cache.Get(url)
	if !hit {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return fmt.Errorf("unexpected status: %d", resp.StatusCode)
		}
		if resp.StatusCode == http.StatusOK {
			jsonData, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			raw = jsonData
		}
		cfg.Cache.Add(url, raw)
	}

	err := json.Unmarshal(raw, &result)
	if err != nil {
		return err
	}

	for _, res := range result.Results {
		fmt.Println(res.Name)
	}

	cfg.Next = result.Next
	cfg.Previous = result.Previous

	return nil
}
