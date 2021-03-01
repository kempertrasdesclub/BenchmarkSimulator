package statisticsBasicsFunctions

func (e *Engine) DefineEventOccurrences(setAllCache, setOne, setSync, invalidateKey, invalidateAll, getAll, getKey int) {
	var amount = float64(setAllCache + setOne + setSync + invalidateKey + invalidateAll + getAll + getKey)

	e.SetAllCache = float64(setAllCache) / amount
	e.SetOne = float64(setOne) / amount
	e.SetSync = float64(setSync) / amount
	e.InvalidateKey = float64(invalidateKey) / amount
	e.InvalidateAll = float64(invalidateAll) / amount
	e.GetAll = float64(getAll) / amount
	e.GetKey = float64(getKey) / amount
}
