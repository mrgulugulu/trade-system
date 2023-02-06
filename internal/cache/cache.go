package cache

import (
	"sync"
	"trade-system/config"

	"github.com/patrickmn/go-cache"
)

var (
	C    cache.Cache
	once sync.Once
)

func init() {
	once.Do(func() {
		C = *cache.New(config.CacheExpirationTime, config.CacheCleanUpInterval)
	})
}
