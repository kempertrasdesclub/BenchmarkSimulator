package cacheAsync

import (
	"cacheSimulator/simulator/data"
	"log"
	"sync"
)

type CacheAsync struct {
	c  map[string]data.Cache
	l  sync.RWMutex
	wg *sync.WaitGroup
}

func (e *CacheAsync) GetFrameworkName() (name string) {
	return "CacheAsync"
}

func (e *CacheAsync) SetAllCache(wg *sync.WaitGroup, content map[string]data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	e.c = make(map[string]data.Cache)
	e.c = content
}

func (e *CacheAsync) Set(wg *sync.WaitGroup, key string, content data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	e.c[key] = content
}

func (e *CacheAsync) InvalidateKey(wg *sync.WaitGroup, key string) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	var found bool
	_, found = e.c[key]
	if found == false {
		log.Println("cacheAsync.InvalidateKey().bug: key not found")
	}

	delete(e.c, key)
}

func (e *CacheAsync) InvalidateAll(wg *sync.WaitGroup) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	e.c = make(map[string]data.Cache)
}

func (e *CacheAsync) GetKey(wg *sync.WaitGroup, key string) (content data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	if e.c[key].UserId != key {
		log.Println("cacheAsync.CacheAsync().error: problema de chave")
	}

	return e.c[key]
}

func (e *CacheAsync) GetAll(wg *sync.WaitGroup) (content map[string]data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	return e.c
}
