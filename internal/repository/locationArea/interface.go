package repository

import "github.com/Glazapolzet/go_pokedex/internal/entity"

type Repository interface {
	GetLocationArea(name string) *entity.LocationArea
}
