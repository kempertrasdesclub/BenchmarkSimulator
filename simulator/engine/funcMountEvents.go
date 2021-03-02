package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
)

func (e *Engine) mountEvents() (err error) {
	var key string
	var dataCache data.DataCache

	for i := 0; i != e.sizeOfEvents; i += 1 {

		randomEvent := e.getEvent()
		key, dataCache, err = e.getCacheByNumericCounter()
		if err != nil {
			return
		}

		if dataCache.UserId == "" {
			panic(errors.New("id clear bug"))
		}

		if key != dataCache.UserId {
			panic(errors.New("bug userID"))
		}

		e.addEvent(key, dataCache, randomEvent)
	}

	return
}
