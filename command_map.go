package main

import (
    "net/http"
    "json"
)

func commandMap(cfg *config) error {
    resp, err := http.Get(cfg.next)
    if err != nil {
	return err
    }

    defer resp.Body.Close()

    return nil
}
