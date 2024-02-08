package implementation

type pokeprinter struct {
	spriteList map[int]string
}

func NewPokeprinter(spriteList map[int]string) *pokeprinter {
	return &pokeprinter{
		spriteList: spriteList,
	}
}

func (p *pokeprinter) GetPokemonSprite(id int) string {
	return p.spriteList[id]
}
