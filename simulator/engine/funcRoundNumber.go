package engine

import "math"

// roundNumber (English): Round a number
//   value (float64): value to be rounded
//   return rounded without decimal places
//
// roundNumber (Português): Arredonda um número
//   value (float64): valor a ser arredondado
//   return número arredondado sem casas decimais
func (e *Engine) roundNumber(value float64) float64 {
	var roundOn = 0.5
	var places = 0.0

	var round float64
	pow := math.Pow(10, places)
	digit := pow * value
	_, div := math.Modf(digit)

	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	return round / pow
}
