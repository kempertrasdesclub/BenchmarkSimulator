package engine

import (
	"cacheSimulator/simulator/statistics"
	"math/rand"
	"time"
)

// getEvent (Português): Escolhe um evento aleatoriamente
//   event: Evento escolhido
func (e *Engine) getEvent() (event statistics.CacheEvent) {
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	var randNumber = randGenerator.Intn(100)
	var percent = float64(randNumber) / 100.0

	return e.getEventByPercent(percent)
}
