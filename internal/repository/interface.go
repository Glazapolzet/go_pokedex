package repository

import (
	locationAreaRepoI "github.com/Glazapolzet/go_pokedex/internal/repository/locationArea"
	locationAreaListRepoI "github.com/Glazapolzet/go_pokedex/internal/repository/locationAreaList"
	pokemonRepoI "github.com/Glazapolzet/go_pokedex/internal/repository/pokemon"
)

type Repository interface {
	locationAreaListRepoI.Repository
	locationAreaRepoI.Repository
	pokemonRepoI.Repository
}
