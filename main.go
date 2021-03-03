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

	eng.SetDataSize(3 * 1000)
	eng.SetEventsSize(100)
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

	err = eng.Init()
	if err != nil {
		log.Fatalf("engine error: %v", err.Error())
	}

	eng.RunSync()
}
