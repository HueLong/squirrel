package cache

import (
	"hbbapi/util/cache/drivers"
	"hbbapi/util/redis"
)

//Cache
//定义缓存的接口
type Cache interface {
	Get(key string) (string, error)
	GetInt(key string) (int, error)
	GetJson(key string, obj interface{}) error
	MultiGetJson(key []string, obj interface{}) error
	Set(key string, data interface{}, expire int, others ...interface{}) error
	SetJson(key string, data interface{}, expire int, others ...interface{}) error
	MultiSetJson(keyValue map[string]interface{}, expire int) error
	Inc(key string) (int, error)
	IncBy(key string, step int) (int, error)
	Dec(key string) (int, error)
	DecBy(key string, step int) (int, error)
	Del(key string) error
	Exp(key string, expire int) error
	Execute(command string, args ...interface{}) (interface{}, error)
}

type Client interface{}

// GetInstance 获取缓存实例
func GetInstance() Cache {
	var client Client
	client = redis.GetOneRedisClient()
	var cache Cache
	driver := drivers.Redis{Client: client}
	cache = driver
	return cache
}

// GetInstanceWithType 手动指定缓存的类型
func GetInstanceWithType(typeName string) Cache {
	var cache Cache
	switch typeName {
	//
	case "redis":
		//cache = new(drivers.Redis)
	case "memcached":
		//cache = new(drivers.Memcached)
	}
	return cache

}
