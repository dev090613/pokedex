package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

func commandExplore(cfg *config, args ...string) error {
	if 1 > len(args) {
		return fmt.Errorf("please pass args")
	}

	url := "https://pokeapi.co/api/v2/location-area"
	area := args[0]

	url = fmt.Sprintf("%s/%s", url, area)

	var result struct {
		Pokemon_encounters []struct{
			Pokemon		struct {
				Name	string `json:"name"`
			} `json:"pokemon"`
		} `json:"pokemon_encounters"`
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

	for _, encounter := range result.Pokemon_encounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
