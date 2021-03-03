package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
	"sync"
	"time"
)

// run (Português): Roda todos os teste de forma síncrona ou assíncrona.
//   synchronous: true para rodar os testes de forma síncrona (espera o teste acabar para chamar outro teste)
func (e *Engine) run(synchronous bool) (err error) {
	var wg sync.WaitGroup
	var fistEventTime time.Duration
	var startTime time.Time
	var endTime time.Duration

	err = e.init()
	if err != nil {
		return
	}

	for _, interactCode := range e.interactions {
		startTime = time.Now()
		interactCode.SetAllCache(&wg, e.cache)
		fistEventTime = time.Since(startTime)

		startTime = time.Now()
		for _, event := range e.eventList {
			switch event.Event {
			case statistics.KSetAllCache:
				if synchronous == true {
					interactCode.SetAllCache(&wg, e.cache)
				} else {
					go interactCode.SetAllCache(&wg, e.cache)
				}

			case statistics.KSet:
				if synchronous == true {
					interactCode.Set(&wg, event.Key, event.DataCache)
				} else {
					go interactCode.Set(&wg, event.Key, event.DataCache)
				}

			case statistics.KInvalidateKey:
				if synchronous == true {
					interactCode.InvalidateKey(&wg, event.Key)
				} else {
					go interactCode.InvalidateKey(&wg, event.Key)
				}

			case statistics.KInvalidateAll:
				if synchronous == true {
					interactCode.InvalidateAll(&wg)
					interactCode.SetAllCache(&wg, e.cache)
				} else {
					go func(wg *sync.WaitGroup, cache map[string]data.DataCache) {
						interactCode.InvalidateAll(wg)
						interactCode.SetAllCache(wg, cache)
					}(&wg, e.cache)
				}

			case statistics.KGetAll:
				if synchronous == true {
					interactCode.GetAll(&wg)
				} else {
					go interactCode.GetAll(&wg)
				}

			case statistics.KGetKey:
				if synchronous == true {
					interactCode.GetKey(&wg, event.Key)
				} else {
					go interactCode.GetKey(&wg, event.Key)
				}
			}
		}

		wg.Wait()
		endTime = time.Since(startTime)
		e.report(fistEventTime, endTime, interactCode.GetFrameworkName())
	}

	return
}
