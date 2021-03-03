package go_cache

import (
	"cacheSimulator/simulator/data"
	"github.com/patrickmn/go-cache"
	"log"
	"sync"
)

type GoCache struct {
	c *cache.Cache
}

func (e *GoCache) SetAllCache(wg *sync.WaitGroup, content map[string]data.DataCache) {
	wg.Add(1)
	defer wg.Done()

	e.c = cache.New(cache.NoExpiration, cache.NoExpiration)
	for k, v := range content {
		e.c.Set(k, v, cache.NoExpiration)
	}
}

func (e *GoCache) Set(wg *sync.WaitGroup, key string, content data.DataCache) {
	wg.Add(1)
	defer wg.Done()

	e.c.Set(key, content, cache.NoExpiration)
}

func (e *GoCache) InvalidateKey(wg *sync.WaitGroup, key string) {
	wg.Add(1)
	defer wg.Done()

	e.c.Delete(key)
}

func (e *GoCache) InvalidateAll(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	e.c = cache.New(cache.NoExpiration, cache.NoExpiration)
}

func (e *GoCache) GetKey(wg *sync.WaitGroup, key string) (content data.DataCache) {
	wg.Add(1)
	defer wg.Done()

	con, _ := e.c.Get(key)
	if con == nil {
		log.Println("go_cache.GetKey().error: key not found")
		return
	}
	return con.(data.DataCache)
}

func (e *GoCache) GetAll(wg *sync.WaitGroup) (content map[string]data.DataCache) {
	wg.Add(1)
	defer wg.Done()

	r := make(map[string]data.DataCache)
	i := e.c.Items()
	for k, v := range i {
		r[k] = v.Object.(data.DataCache)
	}

	return r
}

func (e *GoCache) GetFrameworkName() (name string) {
	return "go-cache"
}
