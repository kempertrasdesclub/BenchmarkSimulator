package cacheAsyncLoop

import (
	"cacheSimulator/simulator/data"
	"log"
	"regexp"
	"sync"
)

type CacheAsyncLoop struct {
	c  map[string]data.Cache
	l  sync.RWMutex
	wg *sync.WaitGroup
}

func (e *CacheAsyncLoop) Init() (err error) {
	return
}

func (e *CacheAsyncLoop) End() (err error) {
	return
}

func (e *CacheAsyncLoop) GetFrameworkName() (name string) {
	return "CacheAsyncLoop"
}

func (e *CacheAsyncLoop) SetAllCache(wg *sync.WaitGroup, content map[string]data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	e.c = make(map[string]data.Cache)
	e.c = content
}

func (e *CacheAsyncLoop) Set(wg *sync.WaitGroup, key string, content data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	e.c[key] = content
}

func (e *CacheAsyncLoop) InvalidateKey(wg *sync.WaitGroup, key string) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	var found bool
	_, found = e.c[key]
	if found == false {
		log.Println("cacheAsync.InvalidateKey().bug: key not found")
	}

	for i := range e.c {
		matches, _ := regexp.MatchString(key, i)
		if matches == true {
			delete(e.c, i)
		}
	}
}

func (e *CacheAsyncLoop) InvalidateAll(wg *sync.WaitGroup) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	for i := range e.c {
		delete(e.c, i)
	}
}

func (e *CacheAsyncLoop) GetKey(wg *sync.WaitGroup, key string) (content data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	if e.c[key].UserId != key {
		log.Println("cacheAsync.CacheAsyncLoop().error: problema de chave")
	}

	return e.c[key]
}

func (e *CacheAsyncLoop) GetAll(wg *sync.WaitGroup) (content map[string]data.Cache) {
	e.l.Lock()
	defer e.l.Unlock()

	wg.Add(1)
	defer wg.Done()

	return e.c
}
