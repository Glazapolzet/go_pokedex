package repository

type Repository interface {
	GetNextLocationAreaList() *LocationAreaList
	GetPrevLocationAreaList() *LocationAreaList
	GetLocationArea(name string) *LocationArea
}
