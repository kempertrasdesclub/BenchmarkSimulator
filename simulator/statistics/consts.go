package statistics

// CacheEvent (Português): Recebe a ação causada pelo usuário.
type CacheEvent int

// String (Português): Transforma a ação do usuário em texto
func (e CacheEvent) String() string {
	return eventAsString[e]
}

const (
	KSetAllCache CacheEvent = iota
	KSet
	KInvalidateKey
	KInvalidateAll
	KGetAll
	KGetKey
)

var eventAsString = [...]string{
	"set all cache",
	"set one",
	"invalidate key",
	"invalidate all",
	"get all",
	"get key",
}
