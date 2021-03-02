package statisticsBasicsFunctions

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/interfaces"
	"cacheSimulator/simulator/statistics"
	"errors"
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
	interactions interfaces.Interactions

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

func (e *Engine) Run(frameworkName string) {
	var wg sync.WaitGroup

	startTime := time.Now()
	go e.interactions.SetAllCache(&wg, e.cache)
	fistEventTime := time.Since(startTime)

	startTime = time.Now()
	for _, event := range e.eventList {
		switch event.Event {
		case statistics.KSetAllCache:
			go e.interactions.SetAllCache(&wg, e.cache)

		case statistics.KSet:
			go e.interactions.Set(&wg, event.Key, event.DataCache)

		case statistics.KInvalidateKey:
			go e.interactions.InvalidateKey(&wg, event.Key)

		case statistics.KInvalidateAll:
			go func() {
				e.interactions.InvalidateAll(&wg)
				e.interactions.SetAllCache(&wg, e.cache)
			}()

		case statistics.KGetAll:
			go e.interactions.GetAll(&wg)

		case statistics.KGetKey:
			go e.interactions.GetKey(&wg, event.Key)

		}
	}

	endTime := time.Since(startTime)
	wg.Wait()

	e.report(fistEventTime, endTime, frameworkName)
}

func (e *Engine) report(firstDataTime, timeDuration time.Duration, frameworkName string) {

}

func (e *Engine) SetInterfaceInteractions(interactions interfaces.Interactions) {
	e.interactions = interactions
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