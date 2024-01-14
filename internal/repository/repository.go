package repository

type urlConfig struct {
	BaseUrl                 string
	LocationAreaListUrl     string
	NextLocationAreaListUrl *string
	PrevLocationAreaListUrl *string
}

var urls = urlConfig{
	BaseUrl:                 "https://pokeapi.co/api/v2/",
	LocationAreaListUrl:     "https://pokeapi.co/api/v2/location-area/",
	NextLocationAreaListUrl: nil,
	PrevLocationAreaListUrl: nil,
}
