package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

type LocationResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
	} `json:"names"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
}

type PokeAPIClient interface {
	getNextLocations(config MapConfig) []LocationResponse
	getPrevLocations(config MapConfig) []LocationResponse
}

type Results struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type MapConfig struct {
	Next string `json:"next"`
	Prev string `json:"previous"`
	Result []Results `json:"results"`
}
func getNextLocations(config MapConfig) (MapConfig, error){
	if config.Next == "" {
		config.Next = POKEAPI_BASE_URL + "/location"
	}

	res, err := http.Get(config.Next)

	if err != nil {
		fmt.Println("HTTP Request Error:", err)
		return MapConfig{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return MapConfig{}, fmt.Errorf("Response code is %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Read Body Error:", err)
		return MapConfig{}, err
	}


	curr_config := MapConfig{}

	err = json.Unmarshal(body, &curr_config)

	if err != nil {
		fmt.Println("Unmarshal Error:", err)
		fmt.Println("Raw JSON Response:", string(body))
		return MapConfig{}, err
	}

	return curr_config, nil
}
