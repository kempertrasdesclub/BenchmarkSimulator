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

	err = eng.RunSync()
	if err != nil {
		log.Fatalf("engine error: %v", err.Error())
	}
}

```

**Output**:
```
Framework name: CacheAsync
Data size: 300000
Events size: 1000
First data load time: 130.068958ms
Execution time: 1.49045525s
Events list:
  set all cache: 11
  set one key: 270
  set invalidate one key: 99
  set invalidate all data: 7
  set invalidate all data: 7
  get all: 235
  get key: 378

Framework name: go-cache
Data size: 300000
Events size: 1000
First data load time: 164.771167ms
Execution time: 26.693374s
Events list:
  set all cache: 11
  set one key: 270
  set invalidate one key: 99
  set invalidate all data: 7
  set invalidate all data: 7
  get all: 235
  get key: 378

Framework name: bigCache
Data size: 300000
Events size: 1000
First data load time: 426.301708ms
Execution time: 2m3.669703166s
Events list:
  set all cache: 11
  set one key: 270
  set invalidate one key: 99
  set invalidate all data: 7
  set invalidate all data: 7
  get all: 235
  get key: 378

Framework name: ristretto
Data size: 300000
Events size: 1000
First data load time: 10.206609042s
Execution time: 3m7.769787959s
Events list:
  set all cache: 11
  set one key: 270
  set invalidate one key: 99
  set invalidate all data: 7
  set invalidate all data: 7
  get all: 235
  get key: 378
```

