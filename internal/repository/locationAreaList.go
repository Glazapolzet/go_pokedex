package repository

import (
	"encoding/json"
	"log"
)

type locationAreaBrief struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaList struct {
	Count    int                 `json:"count"`
	Next     *string             `json:"next"`
	Previous *string             `json:"previous"`
	Results  []locationAreaBrief `json:"results"`
}

func (ll *locationAreaList) Unmarshal(data []byte) {
	err := json.Unmarshal(data, ll)
	if err != nil {
		log.Fatal(err)
	}
}

func getLocationAreaList(url string) *locationAreaList {
	data := get(url)
	locationAreaList := &locationAreaList{}

	unmarshal(locationAreaList, data)

	return locationAreaList
}

func GetNextLocationAreaList() *locationAreaList {
	paginationUrls := urls.locationAreaListPaginationUrls

	locationAreaList := getLocationAreaList(paginationUrls.Current)

	if locationAreaList.Next == nil {
		return nil
	}

	if paginationUrls.Prev == nil && paginationUrls.Next == nil {
		paginationUrls.Update(paginationUrls.Current, locationAreaList.Next, locationAreaList.Previous)

		return locationAreaList
	}

	nextLocationAreaList := getLocationAreaList(*locationAreaList.Next)
	paginationUrls.Update(*locationAreaList.Next, nextLocationAreaList.Next, nextLocationAreaList.Previous)

	return nextLocationAreaList
}

func GetPrevLocationAreaList() *locationAreaList {
	paginationUrls := urls.locationAreaListPaginationUrls

	locationAreaList := getLocationAreaList(paginationUrls.Current)

	if locationAreaList.Previous == nil {
		return nil
	}

	prevLocationAreaList := getLocationAreaList(*locationAreaList.Previous)
	paginationUrls.Update(*locationAreaList.Previous, prevLocationAreaList.Next, prevLocationAreaList.Previous)

	return prevLocationAreaList
}
