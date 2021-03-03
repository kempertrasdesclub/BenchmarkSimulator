package engine

import (
	"cacheSimulator/simulator/statistics"
)

// getEventByPercent (Português): Converte o valor percentual em ação.
//   percent: Valor percentual de ponto flutuante entre 0 e 1
//   event:   Evento selecionado
func (e *Engine) getEventByPercent(percent float64) (event statistics.CacheEvent) {
	if e.setAllCache > percent {
		return statistics.KSetAllCache
	}

	if e.setAllCache+e.setOne > percent {
		return statistics.KSet
	}

	if e.setAllCache+e.setOne+e.invalidateKey > percent {
		return statistics.KInvalidateKey
	}

	if e.setAllCache+e.setOne+e.invalidateKey+e.invalidateAll > percent {
		return statistics.KInvalidateAll
	}

	if e.setAllCache+e.setOne+e.invalidateKey+e.invalidateAll+e.getAll > percent {
		return statistics.KGetAll
	}

	return statistics.KGetKey
}
