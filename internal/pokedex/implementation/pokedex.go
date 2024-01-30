package implementation

import "github.com/Glazapolzet/go_pokedex/internal/repository"

type pokedex struct {
	pokemonList map[string]*repository.Pokemon
}

func NewPokedex() *pokedex {
	return &pokedex{
		pokemonList: make(map[string]*repository.Pokemon),
	}
}

func (p *pokedex) Add(pokemon *repository.Pokemon) {
	p.pokemonList[pokemon.Name] = pokemon
}
