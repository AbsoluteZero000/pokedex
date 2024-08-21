package main

func main() {
	locationsConfig := MapConfig{
		Next: POKEAPI_BASE_URL + "/location",
	}

	startRepl(&locationsConfig)
}
