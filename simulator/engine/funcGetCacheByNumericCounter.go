package engine

import (
	"cacheSimulator/simulator/data"
	"errors"
	"math/rand"
	"time"
)

// getCacheByNumericCounter (Português): Escolhe uma chave aleatoriamente no conjunto de dados
//   key:       valor da chave escolhida
//   dataCache: conteúdo da chave escolhida
//   err:       indicador de erro
func (e *Engine) getCacheByNumericCounter() (key string, dataCache data.DataCache, err error) {
	var randNumber int
	var safeLoop int
	var pass = false
	var found bool
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		safeLoop += 1
		if safeLoop > 100 {
			panic(errors.New("engine.getCacheByNumericCounter().bug: safe loop overflow"))
		}

		counter := 0
		randNumber = randGenerator.Intn(e.sizeOfData - 1)

		for key, dataCache = range e.cache {
			if counter == randNumber {

				pass = true

				_, found = e.doNotRepeatKey[key]
				if found == true {
					pass = false
					break
				}

				e.doNotRepeatKey[key] = true

				return
			}

			counter += 1
		}

		if pass == false {
			continue
		}
	}

	return
}
