package statisticsBasicsFunctions

import (
	"cacheSimulator/simulator/statistics"
)

// getEventByPercent (Português): Converte o valor percentual em ação.
func (e *Engine) getEventByPercent(percent float64) statistics.CacheEvent {
	if e.SetAllCache > percent {
		return statistics.KStatusSetAllCache
	}

	if e.SetAllCache+e.SetOne > percent {
		return statistics.KStatusSet
	}

	if e.SetAllCache+e.SetOne+e.SetSync > percent {
		return statistics.KStatusSetSync
	}

	if e.SetAllCache+e.SetOne+e.SetSync+e.InvalidateKey > percent {
		return statistics.KStatusInvalidateKey
	}

	if e.SetAllCache+e.SetOne+e.SetSync+e.InvalidateKey+e.InvalidateAll > percent {
		return statistics.KStatusInvalidateAll
	}

	if e.SetAllCache+e.SetOne+e.SetSync+e.InvalidateKey+e.InvalidateAll+e.GetAll > percent {
		return statistics.KStatusGetAll
	}

	return statistics.KStatusGetKey
}
