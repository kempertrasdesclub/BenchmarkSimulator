package engine

import (
	"cacheSimulator/simulator/interfaces"
)

func (e *Engine) SetInterfaceData(data interfaces.Data) {
	e.data = data
}
