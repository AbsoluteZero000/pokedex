package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(config *Config) {

	scanner := bufio.NewScanner((bufio.NewReader(os.Stdin)))
	for {
		var commands []string
		fmt.Print("pokedex > ")

		if scanner.Scan() {
			commands = cleanInput(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		if len(commands) == 0 {
			continue
		}

		command := commands[0]

		if cmd, ok := getCommands()[command]; ok {
			config.args = commands[1:]
			if err := cmd.callback(config); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found")
		}
	}

}
