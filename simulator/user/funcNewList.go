package user

import (
	"cacheSimulator/simulator/data"
	"errors"
)

func NewList(
	eventFunctions Interactions,
	statistcsFunctions Statistics,
	numberOfUsers,
	doesNothing,
	setAllCache,
	setOne,
	setSync,
	invalidateKey,
	invalidateAll int,
) (
	cache *map[string]data.Status,
	err error,
) {

	if numberOfUsers == 0 {
		err = errors.New("the user list must have at least one item")
		return
	}

	if statistcsFunctions != nil {
		statistcsFunctions.DefineEventOcurrences(doesNothing, setAllCache, setOne, setSync, invalidateKey, invalidateAll)
	}

	cache = &map[string]data.Status{}
	*cache = make(map[string]data.Status)

	for i := 0; i != numberOfUsers; i += 1 {
		id := NewId()
		status := data.Status{
			UserId:         id,
			Status:         "o rarto roeu a roupa do rei de roma",
			Manual:         false,
			LastActivityAt: 0,
			ActiveChannel:  "channel",
		}
		(*cache)[id] = status
		if eventFunctions != nil {
			eventFunctions.Populate(id, status)
		}
	}

	return
}
