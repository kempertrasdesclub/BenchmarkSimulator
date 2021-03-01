package statistics

// CacheEvent (Português): Recebe a ação causada pelo usuário.
type CacheEvent int

// String (Português): Transforma a ação do usuário em texto
func (e CacheEvent) String() string {
	return eventAsString[e]
}

const (
	KStatusSetAllCache CacheEvent = iota
	KStatusSet
	KStatusSetSync
	KStatusInvalidateKey
	KStatusInvalidateAll
	KStatusGetAll
	KStatusGetKey
)

var eventAsString = [...]string{
	"set all cache",
	"set one",
	"set sync",
	"invalidate key",
	"invalidate all",
	"get all",
	"get key",
}
