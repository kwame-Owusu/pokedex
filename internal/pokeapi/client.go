package pokeapi

import (
	"github.com/kwame-Owusu/pokedex/internal/pokecache"
	"math/rand"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
	rng        *rand.Rand
}

func NewClient(timeout time.Duration, rng *rand.Rand) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(5 * time.Second),
		rng:        rng,
	}
}
