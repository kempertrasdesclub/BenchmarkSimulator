package user

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/statics"
)

type Interactions interface {
	StatusSetAllCache(status map[string]data.Status)
	StatusSet(key string, status data.Status)
	StatusSetSync(key string, status data.Status)
	StatusInvalidate(key string)

	GetCacheCopy() (cache map[string]data.Status)
	Populate(key string, status data.Status)
}

type Statistics interface {
	DefineEventOcurrences(doesNothing, setAllCache, setOne, setSync, invalidateKey, invalidateAll int)
	GetEvent() (action statics.CacheEvent)
}
