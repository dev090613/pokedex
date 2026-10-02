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

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result struct {
		Next *string `json:"next"`
		Previous *string `json:"previous"`
		Results []struct{
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"results"`
	}
	err = json.Unmarshal(jsonData, &result)
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
