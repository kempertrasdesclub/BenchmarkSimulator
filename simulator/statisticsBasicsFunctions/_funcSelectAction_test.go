package statisticsBasicsFunctions

import (
  "cacheSimulator/simulator/statistics"
  "log"
)

func ExampleSelectUserAction_DefineEventOcurrences() {

	// Note: for this test working, the sum of all values must be 99
	var setAllCacheStatisticsRequired = 15
	var setOneStatisticsRequired = 15
	var setSyncStatisticsRequired = 15
	var invalidateKeyStatisticsRequired = 15
	var invalidateAllStatisticsRequired = 15
	var getAllStatisticsRequired = 15
	var getKeyStatisticsRequired = 10

	var setAllCacheCalculedValue = 0.0
	var setOneCalculedValue = 0.0
	var setSyncCalculedValue = 0.0
	var invalidateKeyCalculedValue = 0.0
	var invalidateAllCalculedValue = 0.0
	var getAllCalculedValue = 0.0
	var getKeyCalculedValue = 0.0
	var sumCalculedValue = 0.0

	userAction := SelectUserAction{}
	userAction.DefineEventOccurrences(
		setAllCacheStatisticsRequired,
		setOneStatisticsRequired,
		setSyncStatisticsRequired,
		invalidateKeyStatisticsRequired,
		invalidateAllStatisticsRequired,
		getAllStatisticsRequired,
		getKeyStatisticsRequired,
	)

	for i := 0; i != 100000; i += 1 {
		action := userAction.GetEvent()

		switch action {
		case statistics.KSetAllCache:
			setAllCacheCalculedValue += 1.0
			sumCalculedValue += 1

		case statistics.KSet:
			setOneCalculedValue += 1.0
			sumCalculedValue += 1

		case statistics.KSetSync:
			setSyncCalculedValue += 1.0
			sumCalculedValue += 1

		case statistics.KInvalidateKey:
			invalidateKeyCalculedValue += 1.0
			sumCalculedValue += 1

		case statistics.KInvalidateAll:
			invalidateAllCalculedValue += 1.0
			sumCalculedValue += 1
			
		case statistics.KGetAll:
			getAllCalculedValue += 1.0
			sumCalculedValue += 1

		case statistics.KGetKey:
			getKeyCalculedValue += 1.0
			sumCalculedValue += 1

		default:
			log.Fatal("Error")
		}
	}

	pass := true
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

	if userAction.RoundNumber(getAllCalculedValue/sumCalculedValue*100) != float64(getAllStatisticsRequired) {
		pass = false
	}

	if userAction.RoundNumber(getKeyCalculedValue/sumCalculedValue*100) != float64(getKeyStatisticsRequired) {
		pass = false
	}

	log.Printf("set all cache: %.2f\n", setAllCacheCalculedValue/sumCalculedValue*100)
	log.Printf("set one: %.2f\n", setOneCalculedValue/sumCalculedValue*100)
	log.Printf("set sync: %.2f\n", setSyncCalculedValue/sumCalculedValue*100)
	log.Printf("invalidate key: %.2f\n", invalidateKeyCalculedValue/sumCalculedValue*100)
	log.Printf("invalidate all: %.2f\n", invalidateAllCalculedValue/sumCalculedValue*100)
	log.Printf("getAll: %.2f\n", getAllCalculedValue/sumCalculedValue*100)
	log.Printf("getKey: %.2f\n", getKeyCalculedValue/sumCalculedValue*100)

	if pass == false {
		log.Fatal("error in the estimated percentage value")
	}

	// output:
	//
}
