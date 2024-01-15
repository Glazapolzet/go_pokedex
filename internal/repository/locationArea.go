package repository

import (
	"encoding/json"
	"fmt"
	"log"
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
	urls.LocationAreaListUrl = current
	urls.NextLocationAreaListUrl = next
	urls.PrevLocationAreaListUrl = previous

	return nil
}

func GetNextLocationAreaList() *locationAreaList {
	locationAreaList := getLocationAreaList(urls.LocationAreaListUrl)

	if locationAreaList.Next == nil {
		fmt.Printf("\nNo more pages left..\n\n")

		return nil
	}

	if urls.PrevLocationAreaListUrl == nil && urls.NextLocationAreaListUrl == nil {
		updateLocationAreaListUrls(urls.LocationAreaListUrl, locationAreaList.Next, locationAreaList.Previous)

		return locationAreaList
	}

	nextLocationAreaList := getLocationAreaList(*locationAreaList.Next)
	updateLocationAreaListUrls(*locationAreaList.Next, nextLocationAreaList.Next, nextLocationAreaList.Previous)

	return nextLocationAreaList
}

func GetPrevLocationAreaList() *locationAreaList {
	locationAreaList := getLocationAreaList(urls.LocationAreaListUrl)

	if locationAreaList.Previous == nil {
		fmt.Printf("\nNo more pages left..\n\n")

		return nil
	}

	prevLocationAreaList := getLocationAreaList(*locationAreaList.Previous)
	updateLocationAreaListUrls(*locationAreaList.Previous, prevLocationAreaList.Next, prevLocationAreaList.Previous)

	return prevLocationAreaList
}
