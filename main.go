package main

import (
	"cacheSimulator/projects/cacheAsync"
	"cacheSimulator/projects/ristretto"
	"cacheSimulator/simulator/engine"
	"cacheSimulator/simulator/user"
	"log"
)

func main() {
	var err error
	eng := engine.Engine{}

	eng.SetDataSize(900 * 10)
	eng.SetEventsSize(10 * 100)
	eng.SetInterfaceData(&user.User{})
	eng.DefineEventOccurrences(
		1,
		28,
		10,
		1,
		20,
		40,
	)

	eng.AddInterfaceInteractions(&cacheAsync.CacheAsync{})
	eng.AddInterfaceInteractions(&ristretto.Ristretto{})

	err = eng.RunSync()
	if err != nil {
		log.Fatalf("engine error: %v", err.Error())
	}
}
