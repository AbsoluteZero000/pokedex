package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*MapConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {name: "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "next",
			description: "Explore the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "prev",
			description: "Explore the map backwards",
			callback:    commandMapB,
		},
	}
}

func commandHelp(config *MapConfig) error {
	fmt.Print("Welcome to the PokeDex\n\n")
	fmt.Println("Usage:")
	for k, v := range getCommands() {
		fmt.Printf("\t%s - %s\n", k, v.description)
	}
	return nil
}

func commandExit(config *MapConfig) error {
	fmt.Print("Turning off the PokeDEX...\n\n")
	os.Exit(0)
	return nil
}
func commandMap(config *MapConfig) error {

	locations, err := getNextLocations(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, location := range locations {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapB(config *MapConfig) error {

	locations, err := getPrevLocations(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, location := range locations {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
