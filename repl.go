package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {

	for {
		var command string

		fmt.Print("pokedex > ")
		scanner := bufio.NewScanner((bufio.NewReader(os.Stdin)))

		if scanner.Scan() {
			command = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		if cmd, ok := getCommands()[command]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Println(err)
			}
		}
	}

}
