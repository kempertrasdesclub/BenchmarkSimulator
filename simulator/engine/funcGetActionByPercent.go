package engine

import (
	"cacheSimulator/simulator/statistics"
)

// getEventByPercent (Português): Converte o valor percentual em ação.
func (e *Engine) getEventByPercent(percent float64) statistics.CacheEvent {
	if e.SetAllCache > percent {
		return statistics.KSetAllCache
	}

	if e.SetAllCache+e.SetOne > percent {
		return statistics.KSet
	}

	if e.SetAllCache+e.SetOne+e.InvalidateKey > percent {
		return statistics.KInvalidateKey
	}

	if e.SetAllCache+e.SetOne+e.InvalidateKey+e.InvalidateAll > percent {
		return statistics.KInvalidateAll
	}

	if e.SetAllCache+e.SetOne+e.InvalidateKey+e.InvalidateAll+e.GetAll > percent {
		return statistics.KGetAll
	}

	return statistics.KGetKey
}
