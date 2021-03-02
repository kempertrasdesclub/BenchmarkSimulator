package interfaces

import (
	"cacheSimulator/simulator/data"
	"sync"
)

type Interactions interface {
	SetAllCache(wg *sync.WaitGroup, content map[string]data.DataCache)
	Set(wg *sync.WaitGroup, key string, content data.DataCache)
	InvalidateKey(wg *sync.WaitGroup, key string)
	InvalidateAll(wg *sync.WaitGroup)
	GetKey(wg *sync.WaitGroup, key string) (content data.DataCache)
	GetAll(wg *sync.WaitGroup) (content map[string]data.DataCache)

	GetFrameworkName() (name string)
}
