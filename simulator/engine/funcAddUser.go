package engine

import (
	"cacheSimulator/simulator/data"
)

func (e *Engine) addUser(key string, data data.DataCache) {
	e.cache[key] = data
}
