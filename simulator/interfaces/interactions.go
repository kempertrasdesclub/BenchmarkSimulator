package interfaces

import (
	"cacheSimulator/simulator/data"
	"sync"
)

type Interactions interface {
	SetAllCache(wg *sync.WaitGroup, content map[string]data.DataCache)
	Set(wg *sync.WaitGroup, key string, content data.DataCache)
	SetSync(wg *sync.WaitGroup, key string, content data.DataCache)
	InvalidateKey(wg *sync.WaitGroup, key string)
	InvalidateAll(wg *sync.WaitGroup)
	GetKey(wg *sync.WaitGroup, key string) (content data.DataCache)
	GetAll(wg *sync.WaitGroup) (content []data.DataCache)

	GetCacheCopy(wg *sync.WaitGroup) (cache map[string]data.DataCache)
	Populate(wg *sync.WaitGroup, key string, content data.DataCache)
}
