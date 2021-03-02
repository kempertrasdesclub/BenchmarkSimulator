package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statistics"
)

type Event struct {
	DataCache data.DataCache
	Event     statistics.CacheEvent
	Key       string
}
