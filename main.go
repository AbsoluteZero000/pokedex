package main

func main() {
	locations_url := POKEAPI_BASE_URL + "/location"
	locationsConfig := MapConfig{
		Next: &locations_url,
		Prev: nil,
		Result: nil,
	}

	startRepl(&locationsConfig)
}
