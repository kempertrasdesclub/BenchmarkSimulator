package statisticsBasicsFunctions

// DefineEventOcurrences (Português): Define o valor percentual de ações do usuário entre as opções:
//   doesNothing
//   setAllCache
//   setOne
//   setSync
//   invalidate
func (e *SelectUserAction) DefineEventOcurrences(doesNothing, setAllCache, setOne, setSync, invalidate int) {
	var amount = float64(doesNothing + setAllCache + setOne + setSync + invalidate)

	e.DoesNothing = float64(doesNothing) / amount
	e.SetAllCache = float64(setAllCache) / amount
	e.SetOne = float64(setOne) / amount
	e.SetSync = float64(setSync) / amount
	e.Invalidate = float64(invalidate) / amount
}
