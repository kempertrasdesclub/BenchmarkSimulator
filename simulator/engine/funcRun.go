package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
	"reflect"
	"sync"
	"time"
)

func (e *Engine) mapCopy(dst, src interface{}) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	dv, sv := reflect.ValueOf(dst), reflect.ValueOf(src)

	for _, k := range sv.MapKeys() {
		dv.SetMapIndex(k, sv.MapIndex(k))
	}
}

// run (Português): Roda todos os teste de forma síncrona ou assíncrona.
//   synchronous: true para rodar os testes de forma síncrona (espera o teste acabar para chamar outro teste)
func (e *Engine) run(synchronous bool) (err error) {
	var wg sync.WaitGroup
	var fistEventTime time.Duration
	var startTime time.Time
	var endTime time.Duration
	var cacheCopy = make(map[string]data.DataCache)

	err = e.init()
	if err != nil {
		return
	}

	for _, interactCode := range e.interactions {
		startTime = time.Now()
		e.mapCopy(cacheCopy, e.cache)
		interactCode.SetAllCache(&wg, cacheCopy)
		fistEventTime = time.Since(startTime)

		startTime = time.Now()
		for _, event := range e.eventList {
			switch event.Event {
			case statistics.KSetAllCache:
				if synchronous == true {
					e.mapCopy(cacheCopy, e.cache)
					interactCode.SetAllCache(&wg, cacheCopy)
				} else {
					go func() {
						e.mapCopy(cacheCopy, e.cache)
						interactCode.SetAllCache(&wg, cacheCopy)
					}()
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
					e.mapCopy(cacheCopy, e.cache)
					interactCode.SetAllCache(&wg, cacheCopy)
				} else {
					e.mapCopy(cacheCopy, e.cache)
					go func(wg *sync.WaitGroup, cache map[string]data.DataCache) {
						interactCode.InvalidateAll(wg)
						interactCode.SetAllCache(wg, cache)
					}(&wg, cacheCopy)
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
