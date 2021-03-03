package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
)

// Event (Português): Arquiva os eventos a serem testados
type Event struct {
	DataCache data.Cache
	Event     statistics.CacheEvent
	Key       string
}
