package user

import (
	"cacheSimulator/simulator/data"
)

type User struct{}

func (e *User) NewData() (key string, user data.DataCache) {
	key = e.NewId()
	user = data.DataCache{
		UserId:         key,
		Status:         "o rarto roeu a roupa do rei de roma",
		Manual:         false,
		LastActivityAt: 0,
		ActiveChannel:  "channel",
	}

	return
}
