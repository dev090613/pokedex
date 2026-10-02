package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func cleanInput(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	splited := strings.Fields(input)
	result := make([]string, len(splited))

	for i, s := range splited {
		result[i] = strings.ToLower(s)
	}
	return result
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of previous 20 location areas",
			callback: commandMapb,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if len(input) == 0  {
			continue
		}
		token := cleanInput(input)[0]

		cmd, exists := cfg.commands[token]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scan error:", err)
	}
}

