package engine

import (
	"cacheSimulator/simulator/statistics"
	"math/rand"
	"time"
)

func (e *Engine) GetEvent() (action statistics.CacheEvent) {
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	var randNumber = randGenerator.Intn(100)
	var percent = float64(randNumber) / 100.0

	return e.getEventByPercent(percent)
}
