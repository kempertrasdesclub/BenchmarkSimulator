package interfaces

import (
	"cacheSimulator/simulator/data"
	"sync"
)

// Interactions (Português): Interface com as funções chamadas durante o teste da cache
type Interactions interface {
	// SetAllCache (Português): Popula a cache.
	//     Nota: SetAllCache() é chamado de forma síncrona antes do código rodar a primeira vez.
	//
	//   wg:      Ponteiro de sync.WaitGroup usado para esperar todas as chamadas de evento executarem antes da contagem de
	//            tempo ser calculada.
	//   content: conteúdo completo da cache
	SetAllCache(wg *sync.WaitGroup, content map[string]data.Cache)

	// Set (Português): Define um novo valor para ser adicionado a cache
	//   wg:      Ponteiro de sync.WaitGroup usado para esperar todas as chamadas de evento executarem antes da contagem de
	//            tempo ser calculada.
	//   key:     nome da chave na cache
	//   content: conteúdo da chave
	Set(wg *sync.WaitGroup, key string, content data.Cache)

	// InvalidateKey (Português): Apaga uma chave contida na cache.
	//   wg:  Ponteiro de sync.WaitGroup usado para esperar todas as chamadas de evento executarem antes da contagem de
	//        tempo ser calculada.
	//   key: nome da chave na cache
	InvalidateKey(wg *sync.WaitGroup, key string)

	// InvalidateAll (Português): Apaga toda a cache.
	//     Nota 1: Para fins de testes, InvalidateAll() é sempre seguido de SetAllCache().
	//     Nota 2: Em chamadas assíncronas, SetAllCache() pode não ter terminado a execução quando outras chamadas são
	//             executadas.
	//
	//   wg: Ponteiro de sync.WaitGroup usado para esperar todas as chamadas de evento executarem antes da contagem de
	//       tempo ser calculada.
	InvalidateAll(wg *sync.WaitGroup)

	// GetKey (Português): Recupera um valor contido na cache.
	//   wg:  Ponteiro de sync.WaitGroup usado para esperar todas as chamadas de evento executarem antes da contagem de
	//        tempo ser calculada.
	//   key: nome da chave na cache
	GetKey(wg *sync.WaitGroup, key string) (content data.Cache)

	// GetAll (Português): Recupera o conteúdo completo da cache.
	//   wg:      Ponteiro de sync.WaitGroup usado para esperar todas as chamadas de evento executarem antes da contagem de
	//            tempo ser calculada.
	//   content: Conteúdo completo da cache
	GetAll(wg *sync.WaitGroup) (content map[string]data.Cache)

	// GetFrameworkName (Português): Recupera o nome do framework para o relatório de desempenho.
	//   name: Nome do framework
	GetFrameworkName() (name string)
}
