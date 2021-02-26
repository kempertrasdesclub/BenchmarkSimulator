package statics

// CacheEvent (Português): Recebe a ação causada pelo usuário.
type CacheEvent int

// String (Português): Transforma a ação do usuário em texto
func (e CacheEvent) String() string {
	return eventAsString[e]
}

const (
	KDoesNothing CacheEvent = iota
	KStatusSetAllCache
	KStatusSet
	KStatusSetSync
	KStatusInvalidateKey
	KStatusInvalidateAll
)

var eventAsString = [...]string{
	"does nothing",
	"set all cache",
	"set one",
	"set sync",
	"invalidate key",
	"invalidate all",
}
