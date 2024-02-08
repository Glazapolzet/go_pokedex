package pokedex

import (
	"github.com/Glazapolzet/go_pokedex/internal/entity"
)

type Pokedex interface {
	Add(pokemon *entity.Pokemon)
	GetAll() map[string]*entity.Pokemon
	Get(name string) *entity.Pokemon
}
