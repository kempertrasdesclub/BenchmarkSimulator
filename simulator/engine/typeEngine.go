package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/interfaces"
	"cacheSimulator/simulator/statistics"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Event struct {
	DataCache data.DataCache
	Event     statistics.CacheEvent
	Key       string
}

type Engine struct {
	SizeOfData   int
	SizeOfEvents int

	data         interfaces.Data
	interactions []interfaces.Interactions

	eventList []Event
	cache     map[string]data.DataCache

	totalSetAllCache   int
	totalSetOne        int
	totalInvalidateKey int
	totalInvalidateAll int
	totalGetAll        int
	totalGetKey        int

	SetAllCache   float64
	SetOne        float64
	InvalidateKey float64
	InvalidateAll float64
	GetAll        float64
	GetKey        float64
}

func (e *Engine) Run() {
	var wg sync.WaitGroup
	var fistEventTime time.Duration
	var startTime time.Time
	var endTime time.Duration

	for _, interactCode := range e.interactions {
		startTime = time.Now()
		interactCode.SetAllCache(&wg, e.cache)
		fistEventTime = time.Since(startTime)

		startTime = time.Now()
		for _, event := range e.eventList {
			switch event.Event {
			case statistics.KSetAllCache:
				go interactCode.SetAllCache(&wg, e.cache)

			case statistics.KSet:
				go interactCode.Set(&wg, event.Key, event.DataCache)

			case statistics.KInvalidateKey:
				go interactCode.InvalidateKey(&wg, event.Key)

			case statistics.KInvalidateAll:
				go func() {
					interactCode.InvalidateAll(&wg)
					interactCode.SetAllCache(&wg, e.cache)
				}()

			case statistics.KGetAll:
				go interactCode.GetAll(&wg)

			case statistics.KGetKey:
				go interactCode.GetKey(&wg, event.Key)

			}
		}

		endTime = time.Since(startTime)
		e.report(fistEventTime, endTime, interactCode.GetFrameworkName())
	}
	wg.Wait()

}

func (e *Engine) report(firstDataTime, timeDuration time.Duration, frameworkName string) {
	fmt.Printf("Framework name: %v\n", frameworkName)
	fmt.Printf("First data load time: %v\n", firstDataTime)
	fmt.Printf("Execution time: %v\n", timeDuration)
	fmt.Printf("Events list:\n")
	fmt.Printf("  set all cache: %v\n", e.totalSetAllCache)
	fmt.Printf("  set one key: %v\n", e.totalSetOne)
	fmt.Printf("  set invalidate one key: %v\n", e.totalInvalidateKey)
	fmt.Printf("  set invalidate all data: %v\n", e.totalInvalidateAll)
	fmt.Printf("  set invalidate all data: %v\n", e.totalInvalidateAll)
	fmt.Printf("  get all: %v\n", e.totalGetAll)
	fmt.Printf("  get key: %v\n\n\n", e.totalGetKey)

}

func (e *Engine) AddInterfaceInteractions(interactions interfaces.Interactions) {
	if len(e.interactions) == 0 {
		e.interactions = make([]interfaces.Interactions, 0)
	}

	e.interactions = append(e.interactions, interactions)
}

func (e *Engine) SetInterfaceData(data interfaces.Data) {
	e.data = data
}

func (e *Engine) Init() (err error) {

	if e.data == nil || e.interactions == nil {
		err = errors.New("please, set interfaces first")
		return
	}

	e.eventList = make([]Event, 0)
	e.cache = make(map[string]data.DataCache)

	e.mountData()
	err = e.mountEvents()

	return
}

func (e *Engine) mountEvents() (err error) {
	var key string
	var dataCache data.DataCache
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	var randNumber int

	for i := 0; i != e.SizeOfEvents; i += 1 {
		randNumber = randGenerator.Intn(e.SizeOfData - 1)
		randomEvent := e.GetEvent()
		key, dataCache, err = e.getCacheByNumericCounter(randNumber)
		if err != nil {
			return
		}

		e.addEvent(key, dataCache, randomEvent)
	}

	return
}

func (e *Engine) getCacheByNumericCounter(value int) (key string, dataCache data.DataCache, err error) {
	if value > e.SizeOfData-1 {
		err = errors.New("value out of range")
		return
	}

	counter := 0
	for key, dataCache = range e.cache {
		if counter == value {
			return
		}

		counter += 1
	}

	return
}

func (e *Engine) mountData() {
	for i := 0; i != e.SizeOfData; i += 1 {
		key, keyData := e.data.NewData()
		e.cache[key] = keyData
	}
}

func (e *Engine) addEvent(key string, dataCache data.DataCache, event statistics.CacheEvent) {
	e.eventList = append(e.eventList, Event{
		DataCache: dataCache,
		Event:     event,
		Key:       key,
	})

	switch event {
	case statistics.KSetAllCache:
		e.totalSetAllCache += 1

	case statistics.KSet:
		e.totalSetOne += 1

	case statistics.KInvalidateKey:
		e.totalInvalidateKey += 1

	case statistics.KInvalidateAll:
		e.totalInvalidateAll += 1

	case statistics.KGetAll:
		e.totalGetAll += 1

	case statistics.KGetKey:
		e.totalGetKey += 1
	}
}

func (e *Engine) addUser(key string, data data.DataCache) {
	e.cache[key] = data
}
