package interfaces

import (
	"cacheSimulator/simulator/data"
)

type Interactions interface {
	StatusSetAllCache(status map[string]data.DataCache)
	StatusSet(key string, status data.DataCache)
	StatusSetSync(key string, status data.DataCache)
	StatusInvalidate(key string)

	GetCacheCopy() (cache map[string]data.DataCache)
	Populate(key string, status data.DataCache)
}
