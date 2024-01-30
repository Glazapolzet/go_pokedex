package apiimplementation

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Glazapolzet/go_pokedex/internal/api"
	"github.com/Glazapolzet/go_pokedex/internal/cache"
	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

type repo struct {
	apiCache *cache.Cache
	urls     *PokeApiUrls
}

func NewRepository(cacheTtl time.Duration) *repo {
	r := &repo{
		apiCache: cache.NewCache(cacheTtl),
		urls:     urlConfig,
	}

	return r
}

func (r *repo) GetNextLocationAreaList() *repository.LocationAreaList {
	p := r.urls.LocationAreaListPaginationUrls

	locationAreaList := r.getLocationAreaList(p.Current)

	if locationAreaList.Next == nil {
		return nil
	}

	if p.Previous == nil && p.Next == nil {
		p.Set(p.Current, locationAreaList.Next, locationAreaList.Previous)

		return locationAreaList
	}

	nextLocationAreaList := r.getLocationAreaList(*locationAreaList.Next)

	p.Set(
		*locationAreaList.Next,
		nextLocationAreaList.Next,
		nextLocationAreaList.Previous,
	)

	return nextLocationAreaList
}

func (r *repo) GetPrevLocationAreaList() *repository.LocationAreaList {
	p := r.urls.LocationAreaListPaginationUrls

	locationAreaList := r.getLocationAreaList(p.Current)

	if locationAreaList.Previous == nil {
		return nil
	}

	prevLocationAreaList := r.getLocationAreaList(*locationAreaList.Previous)

	p.Set(
		*locationAreaList.Previous,
		prevLocationAreaList.Next,
		prevLocationAreaList.Previous,
	)

	return prevLocationAreaList
}

func (r *repo) GetLocationArea(name string) *repository.LocationArea {
	return r.getLocationArea(r.urls.LocationAreaListUrl + name)
}

func (r *repo) GetPokemon(name string) *repository.Pokemon {
	return r.getPokemon(r.urls.PokemonUrl + name)
}

func (r *repo) getLocationAreaList(url string) *repository.LocationAreaList {
	data := r.getFromApi(url)

	locationAreaList := repository.NewLocationAreaList()

	err := json.Unmarshal(data, locationAreaList)
	if err != nil {
		log.Fatal(err)
	}

	return locationAreaList
}

func (r *repo) getLocationArea(url string) *repository.LocationArea {
	data := r.getFromApi(url)

	locationArea := repository.NewLocationArea()

	err := json.Unmarshal(data, locationArea)
	if err != nil {
		log.Fatal(err)
	}

	return locationArea
}

func (r *repo) getPokemon(url string) *repository.Pokemon {
	data := r.getFromApi(url)

	pokemon := repository.NewPokemon()

	err := json.Unmarshal(data, pokemon)
	if err != nil {
		log.Fatal(err)
	}

	return pokemon
}

func (r *repo) getFromApi(url string) []byte {
	cachedData, ok := r.apiCache.Get(url)

	if ok {
		return cachedData
	}

	data := api.Get(url)
	r.apiCache.Add(url, data)

	return data
}
