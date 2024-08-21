package main

import "strings"

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
