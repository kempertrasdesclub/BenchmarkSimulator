package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
)

func (e *Engine) addEvent(key string, dataCache data.DataCache, event statistics.CacheEvent) {
	//fixme: não remover ainda.
	//este código impede chaves duplicadas para evitar que delete(cache, key)
	//for _, v := range e.eventList {
	//	if key == "" {
	//		panic(errors.New("engine.addEvent().bug: key is blank"))
	//	}
	//	if key == v.Key {
	//		fmt.Printf("key:   %v\nv.Key: %v\n", key, v.Key)
	//		panic(errors.New("engine.addEvent().bug: duplicated key added"))
	//	}
	//}

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
