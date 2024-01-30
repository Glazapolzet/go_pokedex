package apiimplementation

type PaginationUrls struct {
	Current  string
	Next     *string
	Previous *string
}

func (p *PaginationUrls) Set(current string, next *string, prev *string) {
	p.Current = current
	p.Next = next
	p.Previous = prev
}

type PokeApiUrls struct {
	BaseUrl                        string
	LocationAreaListUrl            string
	LocationAreaListPaginationUrls *PaginationUrls
	PokemonUrl                     string
}

var urlConfig = &PokeApiUrls{
	BaseUrl:             "https://pokeapi.co/api/v2/",
	LocationAreaListUrl: "https://pokeapi.co/api/v2/location-area/",
	LocationAreaListPaginationUrls: &PaginationUrls{
		Current:  "https://pokeapi.co/api/v2/location-area/",
		Next:     nil,
		Previous: nil,
	},
	PokemonUrl: "https://pokeapi.co/api/v2/pokemon/",
}
