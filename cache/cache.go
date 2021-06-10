package cache

import (
	"ethcache/lrucache"
)

var cache = lrucache.New(20)

func Cache() *lrucache.LRUCache {
	return &cache
}
