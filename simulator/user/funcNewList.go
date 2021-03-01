package user

import (
	"cacheSimulator/simulator/data"
)

func NewData() (key string, user data.DataCache) {
	key = NewId()
	user = data.DataCache{
		UserId:         key,
		Status:         "o rarto roeu a roupa do rei de roma",
		Manual:         false,
		LastActivityAt: 0,
		ActiveChannel:  "channel",
	}

	return
}
