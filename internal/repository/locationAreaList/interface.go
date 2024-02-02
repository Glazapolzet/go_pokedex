package repository

import "github.com/Glazapolzet/go_pokedex/internal/entity"

type Repository interface {
	GetLocationAreaListNext() *entity.LocationAreaList
	GetLocationAreaListPrevious() *entity.LocationAreaList
}
