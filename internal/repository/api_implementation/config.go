package apiimplementation

type pokeApiEndpoints struct {
	BaseUrl                        string
	locationAreaListUrl            string
	locationAreaListPaginationUrls *paginationUrls
}

var urlConfig = &pokeApiEndpoints{
	BaseUrl:             "https://pokeapi.co/api/v2/",
	locationAreaListUrl: "https://pokeapi.co/api/v2/location-area/",
	locationAreaListPaginationUrls: &paginationUrls{
		Current:  "https://pokeapi.co/api/v2/location-area/",
		Next:     nil,
		Previous: nil,
	},
}
