package repo

import (
	"context"
	"fmt"
	"kisaanSathi/pkg/config"
	"kisaanSathi/pkg/logger"
	"strings"
	"time"

	rdc "github.com/go-redis/cache/v8"
	rd "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RedisInterface interface {
	//set a value into redis with an expiry
	//	depending on params similar redis command
	//		writeIfNotSet: true -> SET key value EX timeout NX
	//	params
	//		key -> string
	//		value -> interface
	//		timeout -> int | value in millisecond
	//		writeIfNotSet -> bool
	//	interfaces can be directly sent as reference arguments
	//		eg. usage
	//		type Person struct{
	//			Name string,
	//			Age int,
	//		}
	//		var p1 Person{ Name: "John Doe", Age: 25 }
	//		ctx := context.Background()
	//		redisObj := utils.GetRedisObject(ctx)
	//		//ensure timeout is greater than 1s to prevent default timeout and overwrite as needed
	//		redisObj.SetValue("test",p1,1200,true)
	SetValue(ctx context.Context, key string, value interface{}, timeout int, writeIfNotSet bool) error

	//get a value from redis directly into a desired struct
	//	interfaces can be directly sent as reference arguments
	//		eg. usage
	//		type Person struct{
	//			Name string,
	//			Age int,
	//		}
	//		var p1 Person
	//		ctx := context.Background()
	//		redisObj := utils.GetRedisObject(ctx)
	//		redisObj.GetValue("test", &p1)
	GetValue(ctx context.Context, key string, value interface{}) error

	//delete an entry from redis
	DeleteKey(ctx context.Context, key string) error

	//get the TTL for a key set in redis
	GetTTL(ctx context.Context, key string) int

	//check if a key exists in redis
	KeyExists(ctx context.Context, key string) bool

	//add values into redis as a redis hash
	//	accepts a hash key and key-value pairs
	//		eg. usage
	//		type Animal struct{
	//			id		int
	//			name	string
	//			scname 	string
	//			family	string
	//		}
	//		p1 := Person{ name:"elephant", scname:"Loxodonta", family: "mammal"}
	//		dmap := make(map[string]string)
	//		dmap["name"] = p1.name)
	//		dmap["scname"] = p1.scname
	//		dmap["family"] = p1.family
	//		ctx := context.Background()
	//		redisObj := utils.GetRedisObject(ctx)
	//		dkey := fmt.Sprinf("%s_%d", p1.name, p1.id)
	//		redisObj.SetRedisHash(dkey,dmap)
	SetRedisHash(ctx context.Context, key string, kvpairs map[string]string) error

	//fetches individual value for a redis hash key hashmap or all values
	//	all values are fetched if arguments are not provided. only first argument is considered for specific key
	//		eg. usage
	//		type Animal struct{
	//			id		int
	//			name	string
	//			scname 	string
	//			family	string
	//		}
	//		p1 := Person{ name:"elephant", scname:"Loxodonta", family: "mammal"}
	//		dmap := make(map[string]string)
	//		dmap["name"] = p1.name)
	//		dmap["scname"] = p1.scname
	//		dmap["family"] = p1.family
	//		ctx := context.Background()
	//		redisObj := utils.GetRedisObject(ctx)
	//		dkey := fmt.Sprinf("%s_%d", p1.name, p1.id)
	//		dmap, _ := redisObj.GetRedisHashValue(dkey)
	//		dfamily, _:= redisObj.GetRedisHashValue(dkey, "family")
	GetRedisHashValue(ctx context.Context, key string, args ...string) (map[string]string, error)

	DeleteRedisHash(ctx context.Context, key string, fields ...string) error
}

type redisStruct struct {
	Url    string
	Port   int
	Client *rd.Client
	Cache  *rdc.Cache
	//Context context.Context
}

var redisObj *redisStruct = nil

func GetRedisObject(ctx context.Context) (RedisInterface, error) {
	url := config.GetConfig().GetString("repo.redis.host")
	port := config.GetConfig().GetInt("repo.redis.port")
	redisClient := rd.NewClient(&rd.Options{
		Addr:     fmt.Sprintf("%s:%d", url, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//check connection
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	logger.Log(ctx).Info("redis connected: ", zap.String("result", pong))

	redisCache := rdc.New(&rdc.Options{
		Redis: redisClient,
	})

	if redisObj == nil {
		redisObj = &redisStruct{
			Url:    url,
			Port:   port,
			Client: redisClient,
			Cache:  redisCache,
			//Context: ctx,
		}
	}

	return redisObj, nil
}

func CloseRedis(ctx context.Context) error {
	logger.Log(ctx).Info("Closing redis connection START")
	defer logger.Log(ctx).Info("Closing redis connection END")
	if redisObj != nil {
		return redisObj.Client.Close()
	}
	return nil
}

func (obj *redisStruct) SetValue(ctx context.Context, key string, value interface{}, timeout int, writeIfNotSet bool) error {
	cacheItem := &rdc.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Duration(timeout) * time.Millisecond,
		SetNX: writeIfNotSet,
	}
	err := obj.Cache.Set(cacheItem)
	return err
}

func (obj *redisStruct) GetValue(ctx context.Context, key string, value interface{}) error {
	return obj.Cache.Get(ctx, key, &value)
}

func (obj *redisStruct) DeleteKey(ctx context.Context, key string) error {
	return obj.Cache.Delete(ctx, key)
}

func (obj *redisStruct) GetTTL(ctx context.Context, key string) int {
	dur := obj.Client.TTL(ctx, key)
	if dur.Err() != nil {
		return 0
	}
	return int(dur.Val().Milliseconds())
}

func (obj *redisStruct) KeyExists(ctx context.Context, key string) bool {
	return obj.Cache.Exists(ctx, key)
}

func (obj *redisStruct) SetRedisHash(ctx context.Context, key string, kvpairs map[string]string) error {
	if kvpairs == nil {
		return fmt.Errorf("nothing to hash")
	}
	if strings.TrimSpace(key) == "" {
		return fmt.Errorf("key cannot be blank")
	}
	for k, v := range kvpairs {
		val := obj.Client.HSet(ctx, key, k, v)
		if val.Err() != nil {
			return val.Err()
		}
	}
	return nil
}

func (obj *redisStruct) GetRedisHashValue(ctx context.Context, key string, args ...string) (map[string]string, error) {
	if strings.TrimSpace(key) == "" {
		return nil, fmt.Errorf("key cannot be blank")
	}
	field := ""
	if len(args) > 0 {
		field = args[0]
	}
	if strings.TrimSpace(field) == "" {
		kmap := obj.Client.HGetAll(ctx, key)
		rmap, err := kmap.Result()
		if err != nil {
			return nil, err
		}
		return rmap, nil
	} else {
		kval := obj.Client.HGet(ctx, key, field)
		rval, err := kval.Result()
		if err != nil {
			return nil, err
		}
		m := make(map[string]string)
		m[field] = rval
		return m, nil
	}
}

func (obj *redisStruct) DeleteRedisHash(ctx context.Context, key string, fields ...string) error {
	if strings.TrimSpace(key) == "" {
		return fmt.Errorf("key cannot be blank")
	}

	kval := obj.Client.HDel(ctx, key, fields...)
	_, err := kval.Result()
	if err != nil {
		return err
	}
	return nil
}
