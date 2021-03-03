package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
	"math/rand"
	"time"
)

func (e *Engine) getCacheByNumericCounter() (key string, dataCache data.DataCache, err error) {
	var randNumber int
	var safeLoop int
	var pass = false
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		safeLoop += 1
		if safeLoop > 100*1000 {
			panic(errors.New("engine.getCacheByNumericCounter().bug: safe loop overflow"))
		}

		counter := 0
		randNumber = randGenerator.Intn(e.sizeOfData - 1)

		for key, dataCache = range e.cache {
			if counter == randNumber {

				pass = true
				for _, v := range e.eventList {
					if key == "" {
						panic(errors.New("engine.getCacheByNumericCounter().bug: key is blank"))
					}
					if key == v.Key {
						pass = false
						break
					}
				}

				if pass == false {
					break
				}

				return
			}

			counter += 1
		}

		if pass == false {
			continue
		}
	}

	return
}
