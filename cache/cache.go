package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache = cache.New(5*time.Minute, 10*time.Minute)

func Get(key string) interface{} {
	foo, found := Cache.Get(key)
	if found {
		return foo
	}
	return found
}

func Set(key string, emp interface{}) bool {
	Cache.Set(key, emp, cache.NoExpiration)
	return true
}

func GetOption(key string) interface{} {
	opKey := "opt-"+key
	fmt.Println("Looking for cache with key "+opKey)
	dt, exist := Cache.Get(opKey)
	if exist {
		fmt.Println("Cache Key: "+opKey+" was found!")
		return dt
	}

	return nil
}

