// main.go
package main

func main() {
	// endpoint := "location-area"
	// url := fmt.Sprintf("https://pokeapi.co/api/v2/%v/", endpoint)

	cfg := &config{
		commands: getCommand(),
		Next: nil,
		Previous: nil,
	}
	startRepl(cfg)
}
