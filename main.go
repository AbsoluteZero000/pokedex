package main

import (
	"fmt"
	"bufio"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Welcom to the PokeDex")
	return nil
}
func commandExit() error {
	fmt.Println("Turning off the PokeDEX...")
	os.Exit(0)
	return nil
}
func main() {
	cliCommands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	for {
		fmt.Print("pokedex > ")

		var command string

		scanner := bufio.NewScanner((bufio.NewReader(os.Stdin)))

		if scanner.Scan() {
			command = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		if cmd, ok := cliCommands[command]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
