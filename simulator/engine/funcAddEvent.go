package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
)

// addEvent (Português): Adiciona um evento a lista de eventos.
//   A lista de eventos é gerada aleatoriamente e arquiva os eventos de teste do framework na ordem de execução.
//   Para mais informações sobre os eventos, veja interfaces.Interactions.
//
//   key:       chave contida na cache afetada
//   dataCache: novo valor do dado
//   event:     tipo de evento
func (e *Engine) addEvent(key string, dataCache data.Cache, event statistics.CacheEvent) {
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
