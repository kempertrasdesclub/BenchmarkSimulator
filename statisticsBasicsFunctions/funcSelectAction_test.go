package statisticsBasicsFunctions

import (
	"cacheSimulator/statics"
	"log"
)

func ExampleSelectUserAction_DefineEventOcurrences() {

	var doesNotingRequired = 68
	var setAllCacheRequired = 10
	var setOneRequired = 10
	var setSyncRequired = 10
	var invalidateRequired = 2

	var doesNoting = 0.0
	var setAllCache = 0.0
	var setOne = 0.0
	var setSync = 0.0
	var invalidate = 0.0
	var sum = 0.0

	userAction := SelectUserAction{}
	userAction.DefineEventOcurrences(doesNotingRequired, setAllCacheRequired, setOneRequired, setSyncRequired, invalidateRequired)

	for i := 0; i != 100000; i += 1 {
		action := userAction.GetEvent()

		switch action {
		case statics.KDoesNothing:
			doesNoting += 1.0
			sum += 1

		case statics.KStatusSetAllCache:
			setAllCache += 1.0
			sum += 1

		case statics.KStatusSet:
			setOne += 1.0
			sum += 1

		case statics.KStatusSetSync:
			setSync += 1.0
			sum += 1

		case statics.KStatusInvalidate:
			invalidate += 1.0
			sum += 1

		default:
			log.Fatal("Error")
		}
	}

	if userAction.RoundNumber(doesNoting/sum*100) != float64(doesNotingRequired) {
		log.Fatal("error in the estimated percentage value")
	}

	if userAction.RoundNumber(setAllCache/sum*100) != float64(setAllCacheRequired) {
		log.Fatal("error in the estimated percentage value")
	}

	if userAction.RoundNumber(setOne/sum*100) != float64(setOneRequired) {
		log.Fatal("error in the estimated percentage value")
	}

	if userAction.RoundNumber(setSync/sum*100) != float64(setSyncRequired) {
		log.Fatal("error in the estimated percentage value")
	}

	if userAction.RoundNumber(invalidate/sum*100) != float64(invalidateRequired) {
		log.Fatal("error in the estimated percentage value")
	}

	log.Printf("user does noting: %.2f\n", doesNoting/sum*100)
	log.Printf("set all cache: %.2f\n", setAllCache/sum*100)
	log.Printf("set one: %.2f\n", setOne/sum*100)
	log.Printf("set sync: %.2f\n", setSync/sum*100)
	log.Printf("invalidate: %.2f\n", invalidate/sum*100)

	// output:
	//
}
