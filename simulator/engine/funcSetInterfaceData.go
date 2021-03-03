package engine

import (
	"cacheSimulator/simulator/interfaces"
)

// SetInterfaceData (Português): Recebe o ponteiro do objeto responsável pela criação dos dados
func (e *Engine) SetInterfaceData(data interfaces.Data) {
	e.data = data
}
