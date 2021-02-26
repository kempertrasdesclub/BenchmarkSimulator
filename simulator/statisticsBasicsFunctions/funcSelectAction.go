package statisticsBasicsFunctions

// DefineEventOcurrences (Português): Define o valor percentual de ações do usuário entre as opções:
//   doesNothing
//   setAllCache
//   setOne
//   setSync
//   invalidate
func (e *SelectUserAction) DefineEventOcurrences(doesNothing, setAllCache, setOne, setSync, invalidateKey, invalidateAll int) {
	var amount = float64(doesNothing + setAllCache + setOne + setSync + invalidateKey + invalidateAll)

	e.DoesNothing = float64(doesNothing) / amount
	e.SetAllCache = float64(setAllCache) / amount
	e.SetOne = float64(setOne) / amount
	e.SetSync = float64(setSync) / amount
	e.InvalidateKey = float64(invalidateKey) / amount
	e.InvalidateAll = float64(invalidateAll) / amount
}
