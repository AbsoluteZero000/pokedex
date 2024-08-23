package main

import (
	"fmt"
	"os"
	"math/rand"
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
		"explore": {
			name:        "explore",
			description: "Explore the pokemons in a certain location",
			callback:    commandExplore,
		},
		"capture": {
			name: "capture",
			description: "try Catching the pokemon",
			callback: commandCapture,
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
		return err
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
		return err
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

func commandExplore(config *Config) error {
	fmt.Println("Exploring " + config.args[0] + "...")
	pokemons, err := config.exploreLocation(config.args[0])

	if err != nil {
		return err
	}

	fmt.Print(pokemons)

	return nil
}

func commandCapture(conifg *Config) error {
	fmt.Println("Caputring " + conifg.args[0] + "...")

	base_experience, err := conifg.getBaseExperience(conifg.args[0])

	if err != nil {
		fmt.Println(err)
		return err
	}

	randNum := rand.Intn(base_experience)

	if randNum > 40{
		fmt.Println("You caught " + conifg.args[0] + "!")
	} else {
		fmt.Println("You failed to catch " + conifg.args[0] + "!")
	}

	return nil
}
