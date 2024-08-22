package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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

func commandHelp(config *Config) error {
	fmt.Print("Welcome to the PokeDex\n\n")
	fmt.Println("Usage:")
	for k, v := range getCommands() {
		fmt.Printf("\t%s - %s\n", k, v.description)
	}
	return nil
}

func commandExit(config *Config) error {
	fmt.Print("Turning off the PokeDEX...\n\n")
	os.Exit(0)
	return nil
}
func commandMap(config *Config) error {

	locations, err := config.getLocations(config.Next)

	if err != nil {
		if err.Error() == "url is empty" {
			fmt.Println("End of map")
			return nil
		}
		fmt.Println(err)
		return nil
	}

	for _, location := range locations.Result {
		fmt.Printf("%s\n", location.Name)
	}

	config.Next = ""
	config.Prev = ""

	if locations.Next != nil {
		config.Next = *locations.Next
	}

	if locations.Prev != nil {
		config.Prev = *locations.Prev
	}

	return nil
}

func commandMapB(config *Config) error {

	locations, err := config.getLocations(config.Prev)

	if err != nil {
		if err.Error() == "url is empty" {
			fmt.Println("End of map")
			return nil
		}
		fmt.Println(err)
		return nil
	}

	for _, location := range locations.Result {
		fmt.Printf("%s\n", location.Name)
	}

	config.Next = ""
	config.Prev = ""

	if locations.Next != nil {
		config.Next = *locations.Next
	}

	if locations.Prev != nil {
		config.Prev = *locations.Prev
	}

	return nil
}
