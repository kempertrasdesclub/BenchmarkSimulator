package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
)

func (e *Engine) Init() (err error) {

	if e.data == nil || e.interactions == nil {
		err = errors.New("please, set interfaces first")
		return
	}

	e.eventList = make([]Event, 0)
	e.cache = make(map[string]data.DataCache)

	e.mountData()
	err = e.mountEvents()

	return
}
