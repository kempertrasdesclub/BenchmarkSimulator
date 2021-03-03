package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
)

func (e *Engine) Init() (err error) {

	if e.data == nil || e.interactions == nil {
		err = errors.New("engine.Init().error: please, set interfaces first")
		return
	}

	e.cache = make(map[string]data.DataCache)
	e.eventList = make([]Event, 0)

	err = e.mountData()
	if err != nil {
		return
	}

	err = e.mountEvents()

	return
}
