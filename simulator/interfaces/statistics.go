package interfaces

import (
	"cacheSimulator/simulator/statistics"
)

type Statistics interface {
	DefineEventOccurrences(setAllCache, setOne, setSync, invalidateKey, invalidateAll, getAll, getKey int)
	GetEvent() (action statistics.CacheEvent)
}
