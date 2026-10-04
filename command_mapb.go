// command_mapb.go
package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

func commandMapb(cfg *config) error {
	var url string
	if (cfg.Previous == nil) {
		fmt.Println("you're on the first page")
		return nil
	} else { 
		url = *cfg.Previous
	}

	var result struct {
		Next *string `json:"next"`
		Previous *string `json:"previous"`
		Results []struct{
			Name string `json:"name"`
			URL string `json:"url"`
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
