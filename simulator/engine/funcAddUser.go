package engine

import (
	"cacheSimulator/simulator/data"
)

// addDataCache (Português): Adiciona um novo dado ao modelo de cache usado para popular os dados usados no teste
func (e *Engine) addDataCache(key string, dataCache data.DataCache) {
	e.cache[key] = dataCache
}
