package engine

import (
	"cacheSimulator/simulator/interfaces"
)

func (e *Engine) AddInterfaceInteractions(interactions interfaces.Interactions) {
	if len(e.interactions) == 0 {
		e.interactions = make([]interfaces.Interactions, 0)
	}

	e.interactions = append(e.interactions, interactions)
}
