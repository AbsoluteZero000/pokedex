package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
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
		"inspect": {
			name: "inspect",
			description: "get more info about the captured pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "check the pokemons you've captured",
			callback: commandPokedex,
		},
	}
}

func commandHelp(config *Config) error {
	fmt.Print("Welcome to the PokeDex\n\n")
	fmt.Println("Usage:")

	commands := getCommands()
	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("\t%s - %s\n", k, commands[k].description)
	}
	fmt.Println()
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

func commandCapture(config *Config) error {
	fmt.Println("Caputring " + config.args[0] + "...")

	pokemon, err := config.getPokemon(config.args[0])

	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > 40{
		fmt.Println("You caught " + config.args[0] + "!")
		config.capturedPokemon[config.args[0]] = true
	} else {
		fmt.Println("You failed to catch " + config.args[0] + "!")
	}

	return nil
}

func commandInspect(config *Config) error {
	if !config.capturedPokemon[config.args[0]] {
		fmt.Println(config.args[0] + " is not captured")
		return nil
	}

	pokemon, err := config.getPokemon(config.args[0])

	if err != nil {
		return err
	}

	res := "\n"

	res += fmt.Sprintf("Name: %s\n", pokemon.Name)
	res += fmt.Sprintf("Height: %d\n", pokemon.Height)
	res += fmt.Sprintf("Weight: %d\n", pokemon.Weight)
	res += fmt.Sprintf("Base Experience: %d\n", pokemon.BaseExperience)
	res += "Types: \n"
	for _, Type := range pokemon.Types {
		res += fmt.Sprintf("  -%s\n", Type.Type.Name)
	}
	res += "Stats: \n"
	res += fmt.Sprintf("  -hp: %d\n", pokemon.Stats[0].BaseStat)
	res += fmt.Sprintf("  -attack: %d\n", pokemon.Stats[1].BaseStat)
	res += fmt.Sprintf("  -defense: %d\n", pokemon.Stats[2].BaseStat)
	res += fmt.Sprintf("  -special attack: %d\n", pokemon.Stats[3].BaseStat)
	res += fmt.Sprintf("  -special defense: %d\n", pokemon.Stats[4].BaseStat)
	res += fmt.Sprintf("  -speed: %d\n", pokemon.Stats[5].BaseStat)
	fmt.Println(res)
	return nil
}

func commandPokedex(config *Config) error {

	if len(config.capturedPokemon) == 0 {
		fmt.Println("No pokemons captured yet")
		return nil
	}

	res := "Captured pokemons: \n"
	for pokemon := range config.capturedPokemon {
		res += fmt.Sprintf("  - %s\n", pokemon)
	}
	fmt.Println(res)
	return nil
}
