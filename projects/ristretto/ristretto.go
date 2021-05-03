package ristretto

import (
	"cacheSimulator/simulator/data"
	"github.com/dgraph-io/ristretto"
	"sync"
	"time"
)

type Ristretto struct {
	c *ristretto.Cache
}

func (e *Ristretto) Init() (err error) {
	return
}

func (e *Ristretto) End() (err error) {
	return
}

func (e *Ristretto) GetFrameworkName() (name string) {
	return "ristretto"
}

func (e *Ristretto) SetAllCache(wg *sync.WaitGroup, content map[string]data.Cache) {
	wg.Add(1)
	defer wg.Done()

	var err error
	e.c, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	for k, v := range content {
		e.c.Set(k, v, 1)
	}

	time.Sleep(10000 * time.Millisecond)
}

func (e *Ristretto) Set(wg *sync.WaitGroup, key string, content data.Cache) {
	wg.Add(1)
	defer wg.Done()

	e.c.Set(key, content, 1)
	time.Sleep(10 * time.Millisecond)
}

func (e *Ristretto) InvalidateKey(wg *sync.WaitGroup, key string) {
	wg.Add(1)
	defer wg.Done()

	e.c.Del(key)
}

func (e *Ristretto) InvalidateAll(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	e.c.Clear()
}

func (e *Ristretto) GetKey(wg *sync.WaitGroup, key string) (content data.Cache) {
	wg.Add(1)
	defer wg.Done()

	keyContent, _ := e.c.Get(key)

	if keyContent == nil {
		return data.Cache{}
	}

	return keyContent.(data.Cache)
}

func (e *Ristretto) GetAll(wg *sync.WaitGroup) (content map[string]data.Cache) {
	wg.Add(1)
	defer wg.Done()

	// fixme: não tem getAll()

	return nil
}
