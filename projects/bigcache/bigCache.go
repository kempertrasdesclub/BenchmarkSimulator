package bigcache

import (
	"cacheSimulator/simulator/data"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"log"
	"reflect"
	"sync"
	"time"
)

type Bigcache struct {
	c     *bigcache.BigCache
	mutex sync.RWMutex
}

func (e *Bigcache) mapCopy(dst, src interface{}) {
	dv, sv := reflect.ValueOf(dst), reflect.ValueOf(src)

	for _, k := range sv.MapKeys() {
		dv.SetMapIndex(k, sv.MapIndex(k))
	}
}

func (e *Bigcache) Init() (err error) {
	return
}

func (e *Bigcache) End() (err error) {
	return
}

func (e *Bigcache) SetAllCache(wg *sync.WaitGroup, content map[string]data.Cache) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	wg.Add(1)
	defer wg.Done()

	var m = make(map[string]data.Cache)
	e.mapCopy(m, content)

	e.c, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	var dataAsByte []byte
	for k, v := range m {
		dataAsByte, _ = json.Marshal(&v)
		_ = e.c.Set(k, dataAsByte)
	}
}

func (e *Bigcache) Set(wg *sync.WaitGroup, key string, content data.Cache) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	wg.Add(1)
	defer wg.Done()

	var dataAsByte []byte
	dataAsByte, _ = json.Marshal(&content)

	_ = e.c.Set(key, dataAsByte)
}

func (e *Bigcache) InvalidateKey(wg *sync.WaitGroup, key string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	wg.Add(1)
	defer wg.Done()

	var err error
	err = e.c.Delete(key)
	if err != nil {
		log.Printf("bigcache.InvalidateKey().error: %v", err)
	}
}

func (e *Bigcache) InvalidateAll(wg *sync.WaitGroup) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	wg.Add(1)
	defer wg.Done()

	e.c, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
}

func (e *Bigcache) GetKey(wg *sync.WaitGroup, key string) (content data.Cache) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	wg.Add(1)
	defer wg.Done()

	d, err := e.c.Get(key)
	if err != nil {
		log.Printf("bigcache.GetKey().error: %v", err)
		return
	}

	var ret data.Cache
	err = json.Unmarshal(d, &ret)
	if err != nil {
		log.Printf("bigcache.GetKey().error: %v", err)
	}

	return ret
}

func (e *Bigcache) GetAll(wg *sync.WaitGroup) (content map[string]data.Cache) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	wg.Add(1)
	defer wg.Done()

	content = make(map[string]data.Cache)

	i := e.c.Iterator()
	for i.SetNext() {
		info, err := i.Value()
		if err != nil {
			log.Printf("bigcache.GetAll().error: %v", err)
			return
		}

		var ret data.Cache
		err = json.Unmarshal(info.Value(), &ret)
		if err != nil {
			log.Printf("bigcache.GetAll().error: %v", err)
			return
		}

		key := info.Key()

		content[key] = ret
	}

	return
}

func (e *Bigcache) GetFrameworkName() (name string) {
	return "bigCache"
}
