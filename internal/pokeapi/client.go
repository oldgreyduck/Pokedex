package pokeapi

import (
    "net/http"
    "time"
)

type pokeapiClient struct {
    httpClient    http.Client
    cache         map[string][]byte
}

func NewpokeapiClient() *pokeapiClient {
    return &pokeapiClient{
        httpClient:  http.Client{
},
        cache: map[string][]byte,
}
}
