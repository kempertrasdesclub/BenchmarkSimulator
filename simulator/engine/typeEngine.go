package engine

import (
	"cacheSimulator/simulator/data"
	"cacheSimulator/simulator/interfaces"
)

type Engine struct {
	sizeOfData   int
	sizeOfEvents int

	data         interfaces.Data
	interactions []interfaces.Interactions

	eventList []Event
	cache     map[string]data.DataCache

	doNotRepeatKeyList map[int]bool

	totalSetAllCache   int
	totalSetOne        int
	totalInvalidateKey int
	totalInvalidateAll int
	totalGetAll        int
	totalGetKey        int

	setAllCache   float64
	setOne        float64
	invalidateKey float64
	invalidateAll float64
	getAll        float64
	getKey        float64
}
