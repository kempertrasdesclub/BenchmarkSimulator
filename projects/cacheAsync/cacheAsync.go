package main

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statics"
	"cacheSimulator/simulator/statisticsBasicsFunctions"
	"cacheSimulator/simulator/user"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

type CacheAsync struct {
	c map[string]data.Status
	l sync.RWMutex
}

func (e *CacheAsync) SetKey(key string, data data.Status) {
	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = data
}

func (e *CacheAsync) SetAll(data map[string]data.Status) {
	e.l.Lock()
	defer e.l.Unlock()

	e.c = data
}

func (e *CacheAsync) Init() {
	e.c = make(map[string]data.Status)
}

func (e *CacheAsync) StatusSetAllCache(newData map[string]data.Status) {
	e.l.Lock()
	defer e.l.Unlock()

	e.c = newData
}

func (e *CacheAsync) StatusSet(key string, keyData data.Status) {
	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = keyData
}

func (e *CacheAsync) StatusSetSync(key string, keyData data.Status) {
	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = keyData
}

func (e *CacheAsync) StatusInvalidate(key string) {
	e.l.Lock()
	defer e.l.Unlock()

	if key == "all" {
		e.c = make(map[string]data.Status)
	} else {
		delete(e.c, key)
	}
}

func (e *CacheAsync) Populate(key string, keyData data.Status) {
	e.l.Lock()
	defer e.l.Unlock()

	e.c[key] = keyData
}

func (e *CacheAsync) GetCacheCopy() (cache map[string]data.Status) {
	e.l.Lock()
	defer e.l.Unlock()

	cache = make(map[string]data.Status)
	dv, sv := reflect.ValueOf(cache), reflect.ValueOf(e.c)

	for _, k := range sv.MapKeys() {
		dv.SetMapIndex(k, sv.MapIndex(k))
	}

	return
}

func getRandKeyAndValue(numberOfUsers int, cache *map[string]data.Status) (key string, keyData data.Status) {
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

	eventController := &CacheAsync{}
	eventController.Init()

	statistcsController := &statisticsBasicsFunctions.SelectUserAction{}

	numberTotalOfEventsInTests := 1000

	cache, err := user.NewList(eventController, statistcsController, numberOfUsers, doesNothingPercent, setAllCachePercent, setOnePercent, setSyncPercent, invalidateKeyPercent, invalidateAllPercent)

	if err != nil {
		log.Fatalf("NewList error: %v", err)
	}

	start := time.Now()
	for i := 0; i != numberTotalOfEventsInTests; i += 1 {
		c := eventController.GetCacheCopy()

		event := statistcsController.GetEvent()

		switch event {
		case statics.KDoesNothing:

		case statics.KStatusInvalidateKey:
			key, _ := getRandKeyAndValue(numberOfUsers, &c)
			go eventController.StatusInvalidate(key)

		case statics.KStatusInvalidateAll:
			go func(cache *map[string]data.Status) {
				eventController.StatusInvalidate("all")
				eventController.StatusSetAllCache(*cache)
			}(cache)

		case statics.KStatusSet:
			key, value := getRandKeyAndValue(numberOfUsers, &c)
			go eventController.StatusSet(key, value)

		case statics.KStatusSetAllCache:
			go eventController.StatusSetAllCache(*cache)

		case statics.KStatusSetSync:
			key, value := getRandKeyAndValue(numberOfUsers, &c)
			go eventController.StatusSetSync(key, value)

		}
	}
	duration := time.Since(start)
	fmt.Printf("Tempo total: %v", duration)

}
