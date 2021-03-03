package user

import (
	"cacheSimulator/simulator/data"
)

// User (Português): Objeto encarregado de criar um novo dado para popular a cache
type User struct{}

// NewData (Português): Popula o dado a ser arquivado na cache
func (e *User) NewData() (key string, user data.Cache, err error) {
	key, err = e.NewId()
	if err != nil {
		return
	}

	user = data.Cache{
		UserId:         key,
		Status:         "o rarto roeu a roupa do rei de roma",
		Manual:         false,
		LastActivityAt: 0,
		ActiveChannel:  "channel",
	}

	return
}
