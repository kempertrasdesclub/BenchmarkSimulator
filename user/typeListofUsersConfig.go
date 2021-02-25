package user

import (
	"cacheSimulator/data"
	"cacheSimulator/statics"
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
	DefineEventOcurrences(doesNothing, setAllCache, setOne, setSync, invalidate int)
	GetEvent() (action statics.CacheEvent)
}
