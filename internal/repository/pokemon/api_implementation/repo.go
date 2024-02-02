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

func (r *repo) GetPokemon(name string) *entity.Pokemon {
	return r.getPokemon(r.url + name)
}

func (r *repo) getPokemon(url string) *entity.Pokemon {
	data := r.get(url)

	pokemon := entity.NewPokemon()

	err := json.Unmarshal(data, pokemon)
	if err != nil {
		log.Fatal(err)
	}

	return pokemon
}
