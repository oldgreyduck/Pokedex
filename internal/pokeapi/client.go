package pokeapi

import (
	"time"
  "pokedex/internal/pokecache"
  "net/http"
)

type PokeApiClient struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewPokeApiClient(cacheInterval time.Duration) *PokeApiClient {
	return &PokeApiClient{
		httpClient: http.Client{
		Timeout: 10 * time.Second,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

