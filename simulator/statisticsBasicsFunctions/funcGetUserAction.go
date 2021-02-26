package statisticsBasicsFunctions

import (
	"cacheSimulator/simulator/statics"
	"math/rand"
	"time"
)

// GetEvent (Português): Retorna uma ação do usuário baseada nas configurações definidas em
// DefineEventOcurrences()
//
//     Para exemplo de uso, veja a função DefineEventOcurrences()
//
func (e *SelectUserAction) GetEvent() (action statics.CacheEvent) {
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	var randNumber = randGenerator.Intn(99) + 1
	var percent = float64(randNumber) / 100.0

	return e.getEventByPercent(percent)
}
