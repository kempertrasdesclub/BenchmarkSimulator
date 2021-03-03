package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
)

// mountEvents (Português): Monta a lista de eventos na ordem do teste.
//   Nota: As chaves dos testes são únicas para evitar apagar chaves inexistentes.
func (e *Engine) mountEvents() (err error) {
	var key string
	var dataCache data.DataCache

	for i := 0; i != e.sizeOfEvents; i += 1 {

		randomEvent := e.getEvent()
		key, dataCache, err = e.getCacheByNumericCounter()
		if err != nil {
			return
		}

		if dataCache.UserId == "" {
			panic(errors.New("id clear bug"))
		}

		if key != dataCache.UserId {
			panic(errors.New("bug userID"))
		}

		e.addEvent(key, dataCache, randomEvent)
	}

	return
}
