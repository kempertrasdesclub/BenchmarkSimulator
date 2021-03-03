package engine

import (
	"cacheSimulator/simulator/interfaces"
)

// AddInterfaceInteractions (Português): Define o ponteiro contendo o objeto de controle do framework em teste.
//   interactions: Ponteiro de objeto compatível com interfaces.Interactions
func (e *Engine) AddInterfaceInteractions(interactions interfaces.Interactions) {
	if len(e.interactions) == 0 {
		e.interactions = make([]interfaces.Interactions, 0)
	}

	e.interactions = append(e.interactions, interactions)
}
