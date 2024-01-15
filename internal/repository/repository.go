package repository

import (
	"time"

	"github.com/Glazapolzet/go_pokedex/internal/api"
	"github.com/Glazapolzet/go_pokedex/internal/cache"
)

type urlConfig struct {
	BaseUrl                 string
	LocationAreaListUrl     string
	NextLocationAreaListUrl *string
	PrevLocationAreaListUrl *string
}

var urls = urlConfig{
	BaseUrl:                 "https://pokeapi.co/api/v2/",
	LocationAreaListUrl:     "https://pokeapi.co/api/v2/location-area/",
	NextLocationAreaListUrl: nil,
	PrevLocationAreaListUrl: nil,
}

var apiCache = cache.NewCache(time.Minute)

func get(url string) []byte {
	cachedData, ok := apiCache.Get(url)

	if ok {
		return cachedData
	}

	data := api.Get(url)
	apiCache.Add(url, data)

	return data
}
