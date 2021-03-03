package engine

// DefineEventOccurrences (Português): Define o valor percentual de cada teste a ser executado.
//   setAllCache:   Valor usado para calcular o percentual de vezes onde a cache vai ser reescrita do zero.
//   setOne:        Valor usado para calcular o percentual de vezes onde uma chave vai ser alterada.
//   invalidateKey: Valor usado para calcular o percentual de vezes onde uma chave vai ser apagada.
//   invalidateAll: Valor usado para calcular o percentual de vezes onde o conteúdo inteiro vai ser apagado e reescrito
//   getAll:        Valor usado para calcular o percentual de vezes onde o conteúdo vai ser reescrito por completo
//   getKey:        Valor usado para calcular o percentual de vezes onde uma chave vai ser recuperada
func (e *Engine) DefineEventOccurrences(setAllCache, setOne, invalidateKey, invalidateAll, getAll, getKey int) {
	var amount = float64(setAllCache + setOne + invalidateKey + invalidateAll + getAll + getKey)

	e.setAllCache = float64(setAllCache) / amount
	e.setOne = float64(setOne) / amount
	e.invalidateKey = float64(invalidateKey) / amount
	e.invalidateAll = float64(invalidateAll) / amount
	e.getAll = float64(getAll) / amount
	e.getKey = float64(getKey) / amount
}
