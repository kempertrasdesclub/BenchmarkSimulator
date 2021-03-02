# BenchmarkSimulator

```golang
package main

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/engine"
	"cacheSimulator/simulator/user"
	"log"
	"sync"
)

type CacheAsync struct {
	c  map[string]data.DataCache
	l  sync.RWMutex
	wg *sync.WaitGroup
}

func (e *CacheAsync) GetFrameworkName() (name string) {
	return "CacheAsync"
}

func (e *CacheAsync) SetAllCache(wg *sync.WaitGroup, content map[string]data.DataCache) {
	wg.Add(1)
	defer wg.Done()
	
	e.l.Lock()
	defer e.l.Unlock()
	
	e.c = make(map[string]data.DataCache)
	e.c = content
}

func (e *CacheAsync) Set(wg *sync.WaitGroup, key string, content data.DataCache) {
	wg.Add(1)
	defer wg.Done()
	
	e.l.Lock()
	defer e.l.Unlock()
	
	e.c[key] = content
}

func (e *CacheAsync) InvalidateKey(wg *sync.WaitGroup, key string) {
	wg.Add(1)
	defer wg.Done()
	
	e.l.Lock()
	defer e.l.Unlock()
	
	delete(e.c, key)
}

func (e *CacheAsync) InvalidateAll(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	
	e.l.Lock()
	defer e.l.Unlock()
	
	e.c = make(map[string]data.DataCache)
}

func (e *CacheAsync) GetKey(wg *sync.WaitGroup, key string) (content data.DataCache) {
	wg.Add(1)
	defer wg.Done()
	
	e.l.Lock()
	defer e.l.Unlock()
	
	return e.c[key]
}

func (e *CacheAsync) GetAll(wg *sync.WaitGroup) (content map[string]data.DataCache) {
	wg.Add(1)
	defer wg.Done()
	
	e.l.Lock()
	defer e.l.Unlock()
	
	return e.c
}


func main() {
	var err error
	eng := engine.Engine{}

	eng.SetDataSize(10 * 1000)
	eng.SetEventsSize(1000)
	eng.SetInterfaceData(&user.User{})
	eng.DefineEventOccurrences(
		1,
		28,
		10,
		1,
		20,
		40,
	)

	eng.AddInterfaceInteractions(&CacheAsync{})

	err = eng.Init()
	if err != nil {
		log.Fatalf("engine error: %v", err.Error())
	}

	eng.RunSync()
}

```