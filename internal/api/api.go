package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (l locationArea) String() string {
	return l.Name
}

type locationAreaList struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

type urlConfig struct {
	BaseUrl                 string
	LocationAreaListUrl     string
	NextLocationAreaListUrl *string
	PrevLocationAreaListUrl *string
}

var apiUrlConfig = urlConfig{
	BaseUrl:                 "https://pokeapi.co/api/v2/",
	LocationAreaListUrl:     "https://pokeapi.co/api/v2/location-area/",
	NextLocationAreaListUrl: nil,
	PrevLocationAreaListUrl: nil,
}

func get(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	return body
}

func getLocationAreaList(url string) *locationAreaList {
	data := get(url)
	locationAreaList := locationAreaList{}

	err := json.Unmarshal(data, &locationAreaList)
	if err != nil {
		log.Fatal(err)
	}

	return &locationAreaList
}

func updateLocationAreaListUrls(current string, next *string, previous *string) error {
	apiUrlConfig.LocationAreaListUrl = current
	apiUrlConfig.NextLocationAreaListUrl = next
	apiUrlConfig.PrevLocationAreaListUrl = previous

	return nil
}

func GetNextLocationAreaList() *locationAreaList {
	locationAreaList := getLocationAreaList(apiUrlConfig.LocationAreaListUrl)

	if locationAreaList.Next == nil {
		fmt.Printf("\nNo more pages left..\n\n")

		return nil
	}

	if apiUrlConfig.PrevLocationAreaListUrl == nil && apiUrlConfig.NextLocationAreaListUrl == nil {
		updateLocationAreaListUrls(apiUrlConfig.LocationAreaListUrl, locationAreaList.Next, locationAreaList.Previous)

		return locationAreaList
	}

	nextLocationAreaList := getLocationAreaList(*locationAreaList.Next)
	updateLocationAreaListUrls(*locationAreaList.Next, nextLocationAreaList.Next, nextLocationAreaList.Previous)

	return nextLocationAreaList
}

func GetPrevLocationAreaList() *locationAreaList {
	locationAreaList := getLocationAreaList(apiUrlConfig.LocationAreaListUrl)

	if locationAreaList.Previous == nil {
		fmt.Printf("\nNo more pages left..\n\n")

		return nil
	}

	prevLocationAreaList := getLocationAreaList(*locationAreaList.Previous)
	updateLocationAreaListUrls(*locationAreaList.Previous, prevLocationAreaList.Next, prevLocationAreaList.Previous)

	return prevLocationAreaList
}
