package user

import (
	"cacheSimulator/data"
	"cacheSimulator/statics"
	"cacheSimulator/statisticsBasicsFunctions"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sync"
	"testing"
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

type LockUnlockMethod struct {
	Cache map[string]data.Status
	m     sync.RWMutex
}

func (e *LockUnlockMethod) Init() {
	e.Cache = make(map[string]data.Status)
}

func (e *LockUnlockMethod) StatusSetAllCache(newData map[string]data.Status) {
	e.m.Lock()
	defer e.m.Unlock()

	e.Cache = newData
}

func (e *LockUnlockMethod) StatusSet(key string, keyData data.Status) {
	e.m.Lock()
	defer e.m.Unlock()

	e.Cache[key] = keyData
}

func (e *LockUnlockMethod) StatusSetSync(key string, keyData data.Status) {
	e.m.Lock()
	defer e.m.Unlock()

	e.Cache[key] = keyData
}

func (e *LockUnlockMethod) StatusInvalidate(key string) {
	e.m.Lock()
	defer e.m.Unlock()

	if key == "all" {
		e.Cache = make(map[string]data.Status)
	} else {
		delete(e.Cache, key)
	}
}

func (e *LockUnlockMethod) Populate(key string, keyData data.Status) {
	e.m.Lock()
	defer e.m.Unlock()

	e.Cache[key] = keyData
}

func (e *LockUnlockMethod) GetCacheCopy() (cache map[string]data.Status) {
	e.m.Lock()
	defer e.m.Unlock()

	cache = make(map[string]data.Status)
	dv, sv := reflect.ValueOf(cache), reflect.ValueOf(e.Cache)

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

func BenchmarkSimulator(b *testing.B) {}

func ExampleNewList() {

	eventController := &CacheAsync{}
	eventController.Init()

	statistcsController := &statisticsBasicsFunctions.SelectUserAction{}

	numberOfUsers := 100000
	doesNothing := 1
	setAllCache := 1
	setOne := 48
	setSync := 48
	invalidate := 2

	assyncronous := true

	numberTotalOfEventsInTests := 1000

	cache, err := NewList(eventController, statistcsController, numberOfUsers, doesNothing, setAllCache, setOne, setSync, invalidate)

	if err != nil {
		log.Fatalf("NewList error: %v", err)
	}

	start := time.Now()
	for i := 0; i != numberTotalOfEventsInTests; i += 1 {
		c := eventController.GetCacheCopy()

		event := statistcsController.GetEvent()

		switch event {
		case statics.KDoesNothing:

		case statics.KStatusInvalidate:
			key, _ := getRandKeyAndValue(numberOfUsers, &c)

			if assyncronous == true {
				go eventController.StatusInvalidate(key)
			} else {
				eventController.StatusInvalidate(key)
			}

		case statics.KStatusSet:
			key, value := getRandKeyAndValue(numberOfUsers, &c)

			if assyncronous == true {
				go eventController.StatusSet(key, value)
			} else {
				eventController.StatusSet(key, value)
			}

		case statics.KStatusSetAllCache:
			if assyncronous == true {
				go eventController.StatusSetAllCache(*cache)
			} else {
				eventController.StatusSetAllCache(*cache)
			}

		case statics.KStatusSetSync:
			key, value := getRandKeyAndValue(numberOfUsers, &c)
			if assyncronous == true {
				go eventController.StatusSetSync(key, value)
			} else {
				eventController.StatusSetSync(key, value)
			}
		}
	}
	duration := time.Since(start)
	fmt.Printf("Tempo total: %v", duration)

	// Output:
	//
}
