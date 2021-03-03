package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
)

// mountData (Português): Monta dados não repetidos na memória de base para os testes
//   err: indicador de erro
func (e *Engine) mountData() (err error) {
	var key string
	var keyData data.Cache

	var safeLoopOverflow = e.sizeOfData * 2
	var safeLoop int
	for i := 0; i != e.sizeOfData; i += 1 {

		if safeLoop > safeLoopOverflow {
			panic(errors.New("engine.mountData().bug: safe loop overflow"))
		}
		safeLoop += 1

		key, keyData, err = e.data.NewData()
		if err != nil {
			return
		}

		var found bool
		_, found = e.cache[key]
		if found == true {
			i -= 1
			continue
		}

		e.cache[key] = keyData
	}

	return
}
