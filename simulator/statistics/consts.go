package statistics

// CacheEvent (Português): Recebe a ação causada pelo usuário.
type CacheEvent int

// String (Português): Transforma a ação do usuário em texto
func (e CacheEvent) String() string {
	return eventAsString[e]
}

const (
	// KSetAllCache (Português): inicializa a cache com um novo dado
	KSetAllCache CacheEvent = iota

	// KSet (Português): popula uma chave da cache com um novo dado
	KSet

	// KInvalidateKey (Português): apaga uma chave da cache
	KInvalidateKey

	// KInvalidateAll (Português): reinicia a cache sem dados
	KInvalidateAll

	// KGetAll (Português): recupera todos os dados da cache
	KGetAll

	// KGetKey (Português): recupera uma chave da cache
	KGetKey
)

// eventAsString (Português): transforma a constante em string
var eventAsString = [...]string{
	"set all cache",
	"set one",
	"invalidate key",
	"invalidate all",
	"get all",
	"get key",
}
