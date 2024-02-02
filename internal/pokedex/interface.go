package pokedex

import "github.com/Glazapolzet/go_pokedex/internal/repository"

type Pokedex interface {
	Add(*repository.Pokemon)
	GetAll() map[string]*repository.Pokemon
	Get(name string) *repository.Pokemon
}
