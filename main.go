package main

import (
	"cacheSimulator/projects/redis"
	"cacheSimulator/simulator/engine"
	"cacheSimulator/simulator/user"
	"log"
)

func main() {
	var err error
	eng := engine.Engine{}

	eng.SetDataSize(300 * 1000)
	eng.SetEventsSize(1 * 1000)
	eng.SetInterfaceData(&user.User{})
	eng.DefineEventOccurrences(
		1,
		28,
		10,
		1,
		20,
		40,
	)

	eng.AddInterfaceInteractions(&redis.CacheRedis{})
	//eng.AddInterfaceInteractions(&cacheAsync.CacheAsync{})
	//eng.AddInterfaceInteractions(&cacheAsyncLoop.CacheAsyncLoop{})
	//eng.AddInterfaceInteractions(&gocache.GoCache{})
	//eng.AddInterfaceInteractions(&bigcache.Bigcache{})
	//eng.AddInterfaceInteractions(&ristretto.Ristretto{})

	err = eng.RunSync()
	if err != nil {
		log.Fatalf("engine error: %v", err.Error())
	}
}
