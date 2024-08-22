package main

import (
	"encoding/json"
	"fmt"
	"github.com/absolutezero000/pokedex/internal/pokecache"
	"io"
	"net/http"
)

type Config struct {
	client http.Client
	Next   string
	Prev   string
	cache  pokecache.Cache
	args   []string
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type MapResponse struct {
	Next   *string  `json:"next"`
	Prev   *string  `json:"previous"`
	Result []Result `json:"results"`
}

type Location struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Config) getLocations(url string) (MapResponse, error) {

	if url == "" {
		return MapResponse{}, fmt.Errorf("url is empty")
	}

	var mapResp MapResponse

	if res, ok := c.cache.Get(url); ok {

		err := json.Unmarshal(res, &mapResp)

		if err != nil {
			fmt.Println("Unmarshal Error:", err)
			return MapResponse{}, err
		}

		return mapResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("HTTP Request Error:", err)
		return MapResponse{}, err
	}

	res, err := c.client.Do(req)

	if err != nil {
		fmt.Println("HTTP Response Error:", err)
		return MapResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return MapResponse{}, fmt.Errorf("response code is %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Read Body Error:", err)
		return MapResponse{}, err
	}

	err = json.Unmarshal(body, &mapResp)

	if err != nil {
		fmt.Println("Unmarshal Error:", err)
		fmt.Println("Raw JSON Response:", string(body))
		return MapResponse{}, err
	}

	c.cache.Add(url, body)

	return mapResp, nil
}

func (c *Config)exploreLocation(name string) (string, error) {

	if name == "" {
		return "", fmt.Errorf("name is empty")
	}

	var l Location

	if res, ok := c.cache.Get(fmt.Sprintf("%s/location-area/%s", POKEAPI_BASE_URL, name)); ok {
		json.Unmarshal(res, &l)
		result := ""

		for _, pokemon := range l.PokemonEncounters {
			result += fmt.Sprintf("- %s\n", pokemon.Pokemon.Name)
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/location-area/%s", POKEAPI_BASE_URL, name), nil)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	res, err := c.client.Do(req)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("response code is %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Read Body Error:", err)
		return "", err
	}
	json.Unmarshal(body, &l)

	result := ""
	for _, pokemon := range l.PokemonEncounters {
		result += fmt.Sprintf("- %s\n", pokemon.Pokemon.Name)
	}

	c.cache.Add(fmt.Sprintf("%s/location-area/%s", POKEAPI_BASE_URL, name), body)

	return result, nil
}
