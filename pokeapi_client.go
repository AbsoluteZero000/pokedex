package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/absolutezero000/pokedex/internal/pokecache"
)

type Config struct {
	client http.Client
	Next string
	Prev string
	cache pokecache.Cache
}

type Result struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type MapResponse struct {
	Next *string `json:"next"`
	Prev *string `json:"previous"`
	Result []Result `json:"results"`
}

func (c* Config)getLocations(url string) (MapResponse, error){

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
