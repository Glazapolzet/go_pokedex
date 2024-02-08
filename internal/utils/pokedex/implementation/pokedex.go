package implementation

import (
	"github.com/Glazapolzet/go_pokedex/internal/entity"
)

type pokedex struct {
	pokemonList map[string]*entity.Pokemon
}

func NewPokedex() *pokedex {
	return &pokedex{
		pokemonList: make(map[string]*entity.Pokemon),
	}
}

func (p *pokedex) GetAll() map[string]*entity.Pokemon {
	return p.pokemonList
}

func (p *pokedex) Get(name string) *entity.Pokemon {
	pokemon, ok := p.pokemonList[name]

	if !ok {
		return nil
	}

	return pokemon
}

func (p *pokedex) Add(pokemon *entity.Pokemon) {
	p.pokemonList[pokemon.Name] = pokemon
}
