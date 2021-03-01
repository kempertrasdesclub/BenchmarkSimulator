package statisticsBasicsFunctions

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/interfaces"
	"cacheSimulator/simulator/statistics"
	"errors"
	"math/rand"
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
	totalSetSync       int
	totalInvalidateKey int
	totalInvalidateAll int
	totalGetAll        int
	totalGetKey        int

	SetAllCache   float64
	SetOne        float64
	SetSync       float64
	InvalidateKey float64
	InvalidateAll float64
	GetAll        float64
	GetKey        float64
}

func (e *Engine) SetInterfaces(data interfaces.Data, interactions interfaces.Interactions) {
	e.data = data
	e.interactions = interactions
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
	case statistics.KStatusSetAllCache:
		e.totalSetAllCache += 1

	case statistics.KStatusSet:
		e.totalSetOne += 1

	case statistics.KStatusSetSync:
		e.totalSetSync += 1

	case statistics.KStatusInvalidateKey:
		e.totalInvalidateKey += 1

	case statistics.KStatusInvalidateAll:
		e.totalInvalidateAll += 1

	case statistics.KStatusGetAll:
		e.totalGetAll += 1

	case statistics.KStatusGetKey:
		e.totalGetKey += 1
	}
}

func (e *Engine) addUser(key string, data data.DataCache) {
	e.cache[key] = data
}
