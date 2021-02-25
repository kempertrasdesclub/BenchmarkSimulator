package statisticsBasicsFunctions

import "cacheSimulator/statics"

// getEventByPercent (Português): Converte o valor percentual em ação.
func (e *SelectUserAction) getEventByPercent(percent float64) statics.CacheEvent {
	if e.DoesNothing > percent {
		return statics.KDoesNothing
	}

	if e.DoesNothing+e.SetAllCache > percent {
		return statics.KStatusSetAllCache
	}

	if e.DoesNothing+e.SetAllCache+e.SetOne > percent {
		return statics.KStatusSet
	}

	if e.DoesNothing+e.SetAllCache+e.SetOne+e.SetSync > percent {
		return statics.KStatusSetSync
	}

	return statics.KStatusInvalidate
}
