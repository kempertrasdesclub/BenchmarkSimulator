package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
	"math/rand"
	"time"
)

func (e *Engine) getCacheByNumericCounter() (key string, dataCache data.DataCache, err error) {
	var found bool
	var randNumber int

	if e.doNotRepeatKeyList == nil {
		e.doNotRepeatKeyList = make(map[int]bool)
	}

	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		randNumber = randGenerator.Intn(e.sizeOfData - 1)
		_, found = e.doNotRepeatKeyList[randNumber]
		if found == false {
			break
		}
	}

	e.doNotRepeatKeyList[randNumber] = true

	counter := 0
	for key, dataCache = range e.cache {
		if counter == randNumber {
			return
		}

		counter += 1
	}

	err = errors.New("getCacheByNumericCounter().error: key not found")
	return
}
