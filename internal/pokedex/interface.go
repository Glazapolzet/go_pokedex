package pokedex

import "github.com/Glazapolzet/go_pokedex/internal/repository"

type Pokedex interface {
	Add(*repository.Pokemon)
}
