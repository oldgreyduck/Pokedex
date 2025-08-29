package pokeapi

type LocationAreaResp struct {
        Count     int               `json:"count"`
        Next      string            `json:"next"`
        Previous  *string           `json:"previous"`
        Results   []LocationArea    `json:"results"`
}

type LocationArea struct {
        Name    string `json:"name"`
        URL     string `json:"url"`
}

type Location struct {
        Name                string `json:"name"`
        PokemonEncounters   []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
        Pokemon       Pokemon `json:"pokemon"`
}

