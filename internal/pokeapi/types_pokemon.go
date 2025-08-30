package pokeapi

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Types          []PokemonType `json:"types"`
	Stats          []PokemonStat `json:"stats"`
	Abilities      []PokemonAbility `json:"abilities"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Ability  struct {
		Name string `json:"name"`
	} `json:"ability"`
}
