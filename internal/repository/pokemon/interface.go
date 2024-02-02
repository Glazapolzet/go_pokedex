package repository

import "github.com/Glazapolzet/go_pokedex/internal/entity"

type Repository interface {
	GetPokemon(name string) *entity.Pokemon
}
