package main

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
	"cacheSimulator/simulator/statisticsBasicsFunctions"
	"cacheSimulator/simulator/user"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

var wg sync.WaitGroup

type CacheAsync struct {
	c  map[string]data.DataCache
	l  sync.RWMutex
	wg *sync.WaitGroup
}

func (e *CacheAsync) SetKey(key string, data data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = data
}

func (e *CacheAsync) SetAll(data map[string]data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	e.c = data
}

func (e *CacheAsync) Init(wg *sync.WaitGroup) {
	e.wg = wg
	e.c = make(map[string]data.DataCache)
}

func (e *CacheAsync) StatusSetAllCache(newData map[string]data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	e.c = newData
}

func (e *CacheAsync) StatusSet(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = keyData
}

func (e *CacheAsync) StatusSetSync(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = keyData
}

func (e *CacheAsync) StatusInvalidate(key string) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	if key == "all" {
		e.c = make(map[string]data.DataCache)
	} else {
		delete(e.c, key)
	}
}

func (e *CacheAsync) Populate(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = keyData
}

func (e *CacheAsync) GetCacheCopy() (cache map[string]data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	cache = make(map[string]data.DataCache)
	dv, sv := reflect.ValueOf(cache), reflect.ValueOf(e.c)

	for _, k := range sv.MapKeys() {
		dv.SetMapIndex(k, sv.MapIndex(k))
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
	setAllCachePercent := 4
	setOnePercent := 45
	setSyncPercent := 45
	invalidateKeyPercent := 4
	invalidateAllPercent := 4
	getAll := 4
	getKey := 4

	eventController := &CacheAsync{}
	eventController.Init(&wg)

	statistcsController := &statisticsBasicsFunctions.SelectUserAction{}

	numberTotalOfEventsInTests := 1000

	cacheData, err := user.NewList(eventController, statistcsController, numberOfUsers, setAllCachePercent, setOnePercent, setSyncPercent, invalidateKeyPercent, invalidateAllPercent, getAll, getKey)

	if err != nil {
		log.Fatalf("NewList error: %v", err)
	}

	start := time.Now()
	for i := 0; i != numberTotalOfEventsInTests; i += 1 {
		c := eventController.GetCacheCopy()

		event := statistcsController.GetEvent()

		switch event {
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

		case statistics.KStatusGetAll:

		case statistics.KStatusGetKey:
		}
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Tempo total: %v", duration)

}
