package main

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
	"cacheSimulator/simulator/statisticsBasicsFunctions"
	"cacheSimulator/simulator/user"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type GoCache struct {
	c  *cache.Cache
	wg *sync.WaitGroup
}

func (e *GoCache) Init(wg *sync.WaitGroup) {
	e.wg = wg
	e.c = cache.New(cache.NoExpiration, cache.NoExpiration)
}

func (e *GoCache) StatusSetAllCache(newData map[string]data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.c = cache.New(cache.NoExpiration, cache.NoExpiration)

	for k, v := range newData {
		e.c.Set(k, v, cache.NoExpiration)
	}
}

func (e *GoCache) StatusSet(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.c.Set(key, keyData, cache.NoExpiration)
}

func (e *GoCache) StatusSetSync(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.c.Set(key, keyData, cache.NoExpiration)
}

func (e *GoCache) StatusInvalidate(key string) {
	e.wg.Add(1)
	defer e.wg.Done()

	if key == "all" {
		e.c = cache.New(cache.NoExpiration, cache.NoExpiration)
	} else {
		e.c.Delete(key)
	}
}

func (e *GoCache) Populate(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.c.Set(key, keyData, cache.NoExpiration)
}

func (e *GoCache) GetCacheCopy() (cache map[string]data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	items := e.c.Items()
	cache = make(map[string]data.DataCache)

	for k, v := range items {
		cache[k] = v.Object.(data.DataCache)
	}
	return
}

func getRandKeyAndValue(numberOfUsers int, cache *map[string]data.DataCache) (key string, keyData data.DataCache) {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	keyAsNumber := randGenerator.Intn(numberOfUsers - 1)
	counter := 0

	for key, keyData = range *cache {
		if counter != keyAsNumber {
			counter += 1
			continue
		}

		return
	}

	return
}

func main() {

	numberOfUsers := 100 * 1000
	doesNothingPercent := 2
	setAllCachePercent := 4
	setOnePercent := 45
	setSyncPercent := 45
	invalidateKeyPercent := 4
	invalidateAllPercent := 4

	eventController := &GoCache{}
	eventController.Init(&wg)

	statistcsController := &statisticsBasicsFunctions.SelectUserAction{}

	numberTotalOfEventsInTests := 1000

	cacheData, err := user.NewList(eventController, statistcsController, numberOfUsers, doesNothingPercent, setAllCachePercent, setOnePercent, setSyncPercent, invalidateKeyPercent, invalidateAllPercent)

	if err != nil {
		log.Fatalf("NewList error: %v", err)
	}

	start := time.Now()
	for i := 0; i != numberTotalOfEventsInTests; i += 1 {
		c := eventController.GetCacheCopy()

		event := statistcsController.GetEvent()

		switch event {
		case statistics.KDoesNothing:

		case statistics.KStatusInvalidateKey:
			key, _ := getRandKeyAndValue(numberOfUsers, &c)
			go eventController.StatusInvalidate(key)

		case statistics.KStatusInvalidateAll:
			go func(cache *map[string]data.DataCache) {
				eventController.StatusInvalidate("all")
				eventController.StatusSetAllCache(*cache)
			}(cacheData)

		case statistics.KStatusSet:
			key, value := getRandKeyAndValue(numberOfUsers, &c)
			go eventController.StatusSet(key, value)

		case statistics.KStatusSetAllCache:
			go eventController.StatusSetAllCache(*cacheData)

		case statistics.KStatusSetSync:
			key, value := getRandKeyAndValue(numberOfUsers, &c)
			go eventController.StatusSetSync(key, value)

		}
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Tempo total: %v", duration)

}
