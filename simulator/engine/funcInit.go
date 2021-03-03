package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
)

// init (Português): Inicializa o framework
//   err: indicador de erro
func (e *Engine) init() (err error) {

	if e.data == nil || e.interactions == nil {
		err = errors.New("engine.init().error: please, set interfaces first")
		return
	}

	e.doNotRepeatKey = make(map[string]bool)
	e.cache = make(map[string]data.DataCache)
	e.eventList = make([]Event, 0)

	err = e.mountData()
	if err != nil {
		return
	}

	err = e.mountEvents()

	return
}
