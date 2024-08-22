package main

import (
	"net/http"
	"time"
	"github.com/absolutezero000/pokedex/internal/pokecache"
)

func main() {
	locationsConfig := Config{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
		Next: POKEAPI_BASE_URL + "/location-area",
		Prev: "",

		cache: pokecache.NewCache(5 * time.Minute),
	}

	startRepl(&locationsConfig)
}
