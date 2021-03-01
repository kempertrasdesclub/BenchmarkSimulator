package interfaces

import (
	"cacheSimulator/simulator/data"
)

type Data interface {
	NewData() (key string, user data.DataCache)
}
