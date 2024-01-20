package repository

import (
	"time"

	"github.com/Glazapolzet/go_pokedex/internal/api"
	"github.com/Glazapolzet/go_pokedex/internal/cache"
)

type paginationUrls struct {
	Current string
	Next    *string
	Prev    *string
}

func (p *paginationUrls) Update(current string, next *string, prev *string) error {
	urls.locationAreaListPaginationUrls = &paginationUrls{
		Current: current,
		Next:    next,
		Prev:    prev,
	}

	return nil
}

type urlConfig struct {
	BaseUrl                        string
	locationAreaListUrl            string
	locationAreaListPaginationUrls *paginationUrls
}

func makeUrlConfig() *urlConfig {
	var urls = &urlConfig{
		BaseUrl:                        "https://pokeapi.co/api/v2/",
		locationAreaListUrl:            "https://pokeapi.co/api/v2/location-area/",
		locationAreaListPaginationUrls: nil,
	}

	urls.locationAreaListPaginationUrls = &paginationUrls{
		Current: urls.locationAreaListUrl,
		Next:    nil,
		Prev:    nil,
	}

	return urls
}

var apiCache = cache.NewCache(time.Minute)
var urls = makeUrlConfig()

func get(url string) []byte {
	cachedData, ok := apiCache.Get(url)

	if ok {
		return cachedData
	}

	data := api.Get(url)
	apiCache.Add(url, data)

	return data
}

type Unmarshaller interface {
	Unmarshal([]byte)
}

func unmarshal(u Unmarshaller, data []byte) {
	u.Unmarshal(data)
}
