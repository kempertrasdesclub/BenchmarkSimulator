package redis

import (
	"cacheSimulator/simulator/data"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"github.com/go-redis/redis/v8"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
	"log"
	"sync"
	"time"
)

const KPrefix = "cachePrefix__"

func (e *CacheRedis) installAndStart(imageName string) (err error) {
	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()
	var newPort nat.Port
	var mountList []mount.Mount
	var defaultRedisPort nat.Port
	var containerId string

	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	go func(c chan iotmakerdocker.ContainerPullStatusSendToChannel) {
		for {
			select {
			case status := <-c:
				log.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					log.Println("image pull complete!")
					return
				}
			}
		}

	}(*pullStatusChannel)

	// stop and remove containers and garbage collector
	err = e.removeAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		return
	}

	defaultRedisPort, err = nat.NewPort("tcp", "6379")
	if err != nil {
		util.TraceToLog()
		return
	}

	newPort, err = nat.NewPort("tcp", "6379")
	if err != nil {
		util.TraceToLog()
		return
	}

	portMap := nat.PortMap{
		// container port number/protocol [tpc/udp]
		defaultRedisPort: []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: newPort.Port(),
			},
		},
	}

	var config = container.Config{
		OpenStdin:    true,
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Env:          []string{},
		Image:        imageName,
	}

	ml := []iotmakerdocker.Mount{
		{
			MountType:   iotmakerdocker.KVolumeMountTypeBind,
			Source:      "./projects/redis/6379.conf",
			Destination: "/etc/redis/6379.conf",
		},
	}

	// define an external MySQL config file path
	mountList, err = iotmakerdocker.NewVolumeMount(ml)
	if err != nil {
		util.TraceToLog()
		return
	}

	_, _, err = dockerSys.ImagePull(config.Image, pullStatusChannel)
	if err != nil {
		util.TraceToLog()
		return
	}

	containerId, err = dockerSys.ContainerCreateWithConfig(
		&config,
		"container_delete_before_test",
		iotmakerdocker.KRestartPolicyNo,
		portMap,
		mountList,
		nil,
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = dockerSys.ContainerStart(containerId)
	if err != nil {
		util.TraceToLog()
		return
	}

	//ready for connections. Version: '8.0.24'

	switch imageName {
	case "redis:latest":
		_, err = dockerSys.ContainerLogsWaitText(containerId, "Ready to accept connections", log.Writer())
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *CacheRedis) removeAllByNameContains(name string) (err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = dockerSys.RemoveAllByNameContains(name)

	return
}

type CacheRedis struct {
	context    context.Context
	connection *redis.Client
	expiration time.Duration
	wg         *sync.WaitGroup
}

func (e *CacheRedis) Init() (err error) {
	err = e.installAndStart("redis:latest")
	if err != nil {
		util.TraceToLog()
		return
	}

	e.context = context.Background()
	e.connection = redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	e.expiration = time.Minute * 60 * 24
	return
}

func (e *CacheRedis) End() (err error) {
	err = e.removeAllByNameContains("delete")
	return
}

func (e *CacheRedis) GetFrameworkName() (name string) {
	return "CacheAsync"
}

func (e *CacheRedis) SetAllCache(wg *sync.WaitGroup, content map[string]data.Cache) {
	wg.Add(1)
	defer wg.Done()

	for key, value := range content {
		e.connection.Set(e.context, KPrefix+key, value, e.expiration)
	}
}

func (e *CacheRedis) Set(wg *sync.WaitGroup, key string, content data.Cache) {
	wg.Add(1)
	defer wg.Done()

	e.connection.Set(e.context, KPrefix+key, content, e.expiration)
}

func (e *CacheRedis) InvalidateKey(wg *sync.WaitGroup, key string) {
	wg.Add(1)
	defer wg.Done()

	e.connection.Del(e.context, KPrefix+key)
}

func (e *CacheRedis) InvalidateAll(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	var err error
	var keys []string
	var cursor uint64

	c := e.connection.Scan(e.context, cursor, KPrefix+`*`, 0)

	keys, cursor, err = c.Result()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	e.connection.Del(e.context, keys...)
}

func (e *CacheRedis) GetKey(wg *sync.WaitGroup, key string) (content data.Cache) {
	var stringCmd *redis.StringCmd
	wg.Add(1)
	defer wg.Done()

	var err error

	stringCmd = e.connection.Get(e.context, KPrefix+key)

	err = stringCmd.Scan(&content)
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	return
}

func (e *CacheRedis) GetAll(wg *sync.WaitGroup) (content map[string]data.Cache) {
	wg.Add(1)
	defer wg.Done()

	var err error
	var keys []string
	var cursor uint64
	var stringCmd *redis.StringCmd
	var dataToConvert data.Cache

	content = make(map[string]data.Cache)

	c := e.connection.Scan(e.context, cursor, KPrefix+`*`, 0)

	keys, cursor, err = c.Result()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	for _, k := range keys {
		stringCmd = e.connection.Get(e.context, k)
		if stringCmd.Err() != nil {
			util.TraceToLog()
			panic(stringCmd.Err())
		}

		err = stringCmd.Scan(&dataToConvert)
		if err != nil {
			util.TraceToLog()
			panic(err)
		}

		content[k] = dataToConvert
	}

	return
}
