package apiimplementation

import (
	"time"

	"github.com/Glazapolzet/go_pokedex/internal/api"
	"github.com/Glazapolzet/go_pokedex/internal/cache"
	"github.com/Glazapolzet/go_pokedex/internal/entity"
	locationAreaRepoI "github.com/Glazapolzet/go_pokedex/internal/repository/locationArea"
	locationAreaApiImplementation "github.com/Glazapolzet/go_pokedex/internal/repository/locationArea/api_implementation"
	locationAreaListRepoI "github.com/Glazapolzet/go_pokedex/internal/repository/locationAreaList"
	locationAreaListApiImplementation "github.com/Glazapolzet/go_pokedex/internal/repository/locationAreaList/api_implementation"
	pokemonRepoI "github.com/Glazapolzet/go_pokedex/internal/repository/pokemon"
	pokemonApiImplementation "github.com/Glazapolzet/go_pokedex/internal/repository/pokemon/api_implementation"
)

type PokeApiUrls struct {
	Base         string
	LocationArea string
	Pokemon      string
}

type repo struct {
	cache                *cache.Cache
	locationAreaRepo     locationAreaRepoI.Repository
	locationAreaListRepo locationAreaListRepoI.Repository
	pokemonRepo          pokemonRepoI.Repository
}

func NewRepo(urls *PokeApiUrls, cacheTTL time.Duration) *repo {
	repo := &repo{
		cache:                cache.NewCache(cacheTTL),
		locationAreaRepo:     nil,
		locationAreaListRepo: nil,
		pokemonRepo:          nil,
	}

	repo.locationAreaRepo = locationAreaApiImplementation.NewRepo(urls.LocationArea, repo.get)
	repo.locationAreaListRepo = locationAreaListApiImplementation.NewRepo(urls.LocationArea, repo.get)
	repo.pokemonRepo = pokemonApiImplementation.NewRepo(urls.Pokemon, repo.get)

	return repo
}

func (r *repo) GetLocationArea(name string) *entity.LocationArea {
	return r.locationAreaRepo.GetLocationArea(name)
}

func (r *repo) GetLocationAreaListNext() *entity.LocationAreaList {
	return r.locationAreaListRepo.GetLocationAreaListNext()
}

func (r *repo) GetLocationAreaListPrevious() *entity.LocationAreaList {
	return r.locationAreaListRepo.GetLocationAreaListPrevious()
}

func (r *repo) GetPokemon(name string) *entity.Pokemon {
	return r.pokemonRepo.GetPokemon(name)
}

func (r *repo) get(url string) []byte {
	cachedData, ok := r.cache.Get(url)

	if ok {
		return cachedData
	}

	data := api.Get(url)
	r.cache.Add(url, data)

	return data
}
