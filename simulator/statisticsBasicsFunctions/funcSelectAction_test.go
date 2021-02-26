package statisticsBasicsFunctions

import (
	"cacheSimulator/simulator/statics"
	"log"
)

func ExampleSelectUserAction_DefineEventOcurrences() {

	// Note: for this test working, the sum of all values must be 100
	var doesNotingStatisticsRequired = 66
	var setAllCacheStatisticsRequired = 10
	var setOneStatisticsRequired = 10
	var setSyncStatisticsRequired = 10
	var invalidateKeyStatisticsRequired = 2
	var invalidateAllStatisticsRequired = 2

	var doesNotingCalculedValue = 0.0
	var setAllCacheCalculedValue = 0.0
	var setOneCalculedValue = 0.0
	var setSyncCalculedValue = 0.0
	var invalidateKeyCalculedValue = 0.0
	var invalidateAllCalculedValue = 0.0
	var sumCalculedValue = 0.0

	userAction := SelectUserAction{}
	userAction.DefineEventOcurrences(
		doesNotingStatisticsRequired,
		setAllCacheStatisticsRequired,
		setOneStatisticsRequired,
		setSyncStatisticsRequired,
		invalidateKeyStatisticsRequired,
		invalidateAllStatisticsRequired,
	)

	for i := 0; i != 100000; i += 1 {
		action := userAction.GetEvent()

		switch action {
		case statics.KDoesNothing:
			doesNotingCalculedValue += 1.0
			sumCalculedValue += 1

		case statics.KStatusSetAllCache:
			setAllCacheCalculedValue += 1.0
			sumCalculedValue += 1

		case statics.KStatusSet:
			setOneCalculedValue += 1.0
			sumCalculedValue += 1

		case statics.KStatusSetSync:
			setSyncCalculedValue += 1.0
			sumCalculedValue += 1

		case statics.KStatusInvalidateKey:
			invalidateKeyCalculedValue += 1.0
			sumCalculedValue += 1

		case statics.KStatusInvalidateAll:
			invalidateAllCalculedValue += 1.0
			sumCalculedValue += 1

		default:
			log.Fatal("Error")
		}
	}

	pass := true
	if userAction.RoundNumber(doesNotingCalculedValue/sumCalculedValue*100) != float64(doesNotingStatisticsRequired) {
		pass = false
	}

	if userAction.RoundNumber(setAllCacheCalculedValue/sumCalculedValue*100) != float64(setAllCacheStatisticsRequired) {
		pass = false
	}

	if userAction.RoundNumber(setOneCalculedValue/sumCalculedValue*100) != float64(setOneStatisticsRequired) {
		pass = false
	}

	if userAction.RoundNumber(setSyncCalculedValue/sumCalculedValue*100) != float64(setSyncStatisticsRequired) {
		pass = false
	}

	if userAction.RoundNumber(invalidateKeyCalculedValue/sumCalculedValue*100) != float64(invalidateKeyStatisticsRequired) {
		pass = false
	}

	if userAction.RoundNumber(invalidateAllCalculedValue/sumCalculedValue*100) != float64(invalidateAllStatisticsRequired) {
		pass = false
	}

	log.Printf("user does noting: %.2f\n", doesNotingCalculedValue/sumCalculedValue*100)
	log.Printf("set all cache: %.2f\n", setAllCacheCalculedValue/sumCalculedValue*100)
	log.Printf("set one: %.2f\n", setOneCalculedValue/sumCalculedValue*100)
	log.Printf("set sync: %.2f\n", setSyncCalculedValue/sumCalculedValue*100)
	log.Printf("invalidate: %.2f\n", invalidateKeyCalculedValue/sumCalculedValue*100)
	log.Printf("invalidate: %.2f\n", invalidateAllCalculedValue/sumCalculedValue*100)

	if pass == false {
		log.Fatal("error in the estimated percentage value")
	}

	// output:
	//
}
