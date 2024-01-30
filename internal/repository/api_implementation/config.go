package apiimplementation

type pokeApiEndpoints struct {
	baseUrl                        string
	locationAreaListUrl            string
	locationAreaListPaginationUrls *paginationUrls
	pokemonUrl                     string
}

var urlConfig = &pokeApiEndpoints{
	baseUrl:             "https://pokeapi.co/api/v2/",
	locationAreaListUrl: "https://pokeapi.co/api/v2/location-area/",
	locationAreaListPaginationUrls: &paginationUrls{
		Current:  "https://pokeapi.co/api/v2/location-area/",
		Next:     nil,
		Previous: nil,
	},
	pokemonUrl: "https://pokeapi.co/api/v2/pokemon/",
}
