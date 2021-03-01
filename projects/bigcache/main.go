package main

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
	"cacheSimulator/simulator/statisticsBasicsFunctions"
	"cacheSimulator/simulator/user"
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"log"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Bigcache struct {
	c  *bigcache.BigCache
	l  sync.RWMutex
	wg *sync.WaitGroup
}

func (e *Bigcache) SetKey(key string, data data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	var dataAsByte []byte
	dataAsByte, _ = json.Marshal(&data)

	_ = e.c.Set(key, dataAsByte)
}

func (e *Bigcache) SetAll(data map[string]data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	var dataAsByte []byte

	for k, v := range data {
		dataAsByte, _ = json.Marshal(&v)
		_ = e.c.Set(k, dataAsByte)
	}
}

func (e *Bigcache) Init(wg *sync.WaitGroup) {
	e.wg = wg
	e.c, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
}

func (e *Bigcache) StatusSetAllCache(newData map[string]data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	var dataAsByte []byte
	for k, v := range newData {
		dataAsByte, _ = json.Marshal(&v)
		_ = e.c.Set(k, dataAsByte)
	}
}

func (e *Bigcache) StatusSet(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	var dataAsByte []byte
	dataAsByte, _ = json.Marshal(&keyData)

	_ = e.c.Set(key, dataAsByte)
}

func (e *Bigcache) StatusSetSync(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	var dataAsByte []byte
	dataAsByte, _ = json.Marshal(&keyData)

	_ = e.c.Set(key, dataAsByte)
}

func (e *Bigcache) StatusInvalidate(key string) {
	e.wg.Add(1)
	defer e.wg.Done()

	e.l.Lock()
	defer e.l.Unlock()

	if key == "all" {
		e.c, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	} else {
		_ = e.c.Delete(key)
	}
}

func (e *Bigcache) Populate(key string, keyData data.DataCache) {
	e.wg.Add(1)
	defer e.wg.Done()

	var dataAsByte []byte
	dataAsByte, _ = json.Marshal(&keyData)

	_ = e.c.Set(key, dataAsByte)
}

func (e *Bigcache) GetCacheCopy() (cache map[string]data.DataCache) {
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
	doesNothingPercent := 2
	setAllCachePercent := 4
	setOnePercent := 15
	setSyncPercent := 15
	invalidateKeyPercent := 4
	invalidateAllPercent := 1
	//get all e get key 25

	eventController := &Bigcache{}
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
