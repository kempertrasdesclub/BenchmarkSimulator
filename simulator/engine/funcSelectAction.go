package engine

func (e *Engine) DefineEventOccurrences(setAllCache, setOne, invalidateKey, invalidateAll, getAll, getKey int) {
	var amount = float64(setAllCache + setOne + invalidateKey + invalidateAll + getAll + getKey)

	e.setAllCache = float64(setAllCache) / amount
	e.setOne = float64(setOne) / amount
	e.invalidateKey = float64(invalidateKey) / amount
	e.invalidateAll = float64(invalidateAll) / amount
	e.getAll = float64(getAll) / amount
	e.getKey = float64(getKey) / amount
}
