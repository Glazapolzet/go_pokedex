package apiimplementation

import (
	"encoding/json"
	"log"

	"github.com/Glazapolzet/go_pokedex/internal/entity"
)

type repo struct {
	url string
	get func(url string) []byte
}

func NewRepo(url string, get func(url string) []byte) *repo {
	return &repo{
		url: url,
		get: get,
	}
}

func (r *repo) GetLocationArea(name string) *entity.LocationArea {
	return r.getLocationArea(r.url + name)
}

func (r *repo) getLocationArea(url string) *entity.LocationArea {
	data := r.get(url)

	locationArea := entity.NewLocationArea()

	err := json.Unmarshal(data, locationArea)
	if err != nil {
		log.Fatal(err)
	}

	return locationArea
}
