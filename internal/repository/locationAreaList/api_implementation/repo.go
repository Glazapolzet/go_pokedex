package apiimplementation

import (
	"encoding/json"
	"log"

	"github.com/Glazapolzet/go_pokedex/internal/entity"
)

type paginationUrls struct {
	Current  string
	Next     *string
	Previous *string
}

type repo struct {
	url            string
	get            func(url string) []byte
	paginationUrls *paginationUrls
}

func NewRepo(url string, get func(url string) []byte) *repo {
	return &repo{
		url: url,
		get: get,
		paginationUrls: &paginationUrls{
			Current:  url,
			Next:     nil,
			Previous: nil,
		},
	}
}

func (r *repo) GetLocationAreaListNext() *entity.LocationAreaList {
	localList := r.paginationUrls
	apiList := r.getLocationAreaList(localList.Current)

	if apiList.Next == nil {
		return nil
	}

	//if the map command was used for the first time
	if localList.Previous == nil && localList.Next == nil {
		r.setPaginationUrls(
			localList.Current,
			apiList.Next,
			apiList.Previous,
		)

		return apiList
	}

	nextApiList := r.getLocationAreaList(*apiList.Next)

	r.setPaginationUrls(
		*apiList.Next,
		nextApiList.Next,
		nextApiList.Previous,
	)

	return nextApiList
}

func (r *repo) GetLocationAreaListPrevious() *entity.LocationAreaList {
	localList := r.paginationUrls
	apiList := r.getLocationAreaList(localList.Current)

	if apiList.Previous == nil {
		return nil
	}

	previousApiList := r.getLocationAreaList(*apiList.Previous)

	r.setPaginationUrls(
		*apiList.Previous,
		previousApiList.Next,
		previousApiList.Previous,
	)

	return previousApiList
}

func (r *repo) getLocationAreaList(url string) *entity.LocationAreaList {
	data := r.get(url)

	locationAreaList := entity.NewLocationAreaList()

	err := json.Unmarshal(data, locationAreaList)
	if err != nil {
		log.Fatal(err)
	}

	return locationAreaList
}

func (r *repo) setPaginationUrls(current string, next *string, previous *string) {
	r.paginationUrls.Current = current
	r.paginationUrls.Next = next
	r.paginationUrls.Previous = previous
}
