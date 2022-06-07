package models

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func SetupCache() *cache.Cache {
	// Create a cache with a default expiration time of 24 hour, and which
	// purges expired items every 10 minutes
	return cache.New(24*60*time.Minute, 10*time.Minute)
}
