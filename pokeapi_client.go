package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

type Result struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type MapConfig struct {
	Next *string `json:"next"`
	Prev *string `json:"previous"`
	Result []Result `json:"results"`
}
func getNextLocations(config *MapConfig) ([]Result, error){
	if config.Next == nil {
		return nil, errors.New("you have reached the end of the map")
	}
	res, err := http.Get(*config.Next)

	if err != nil {
		fmt.Println("HTTP Request Error:", err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response code is %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Read Body Error:", err)
		return nil, err
	}

	err = json.Unmarshal(body, &config)

	if err != nil {
		fmt.Println("Unmarshal Error:", err)
		fmt.Println("Raw JSON Response:", string(body))
		return nil, err
	}

	return config.Result, nil
}


func getPrevLocations(config *MapConfig) ([]Result, error){
	if config.Prev == nil {
		return nil, errors.New("you have reached the start of the map")
	}

	res, err := http.Get(*config.Prev)

	if err != nil {
		fmt.Println("HTTP Request Error:", err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response code is %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Read Body Error:", err)
		return nil, err
	}

	err = json.Unmarshal(body, &config)

	if err != nil {
		fmt.Println("Unmarshal Error:", err)
		fmt.Println("Raw JSON Response:", string(body))
		return nil, err
	}

	return config.Result, nil
}
