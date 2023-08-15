package redis

import (
	"fmt"
	"gin-pro/app/core/system"
	"gin-pro/app/global/consts"
	"gin-pro/library/config"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var redisPool *redis.Pool

func initRedisClientPool() *redis.Pool {
	if redisPool != nil {
		return redisPool
	}
	configYml := config.NewConfig()
	redisPool = &redis.Pool{
		MaxIdle:     configYml.GetInt("Redis.MaxIdle"),                        //最大空闲数
		MaxActive:   configYml.GetInt("Redis.MaxActive"),                      //最大活跃数
		IdleTimeout: configYml.GetDuration("Redis.IdleTimeout") * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.Dial("tcp", configYml.GetString("Redis.Host")+":"+configYml.GetString("Redis.Port"))
			if err != nil {
				system.ZapLog.Error(consts.ErrorsRedisInitConnFail + err.Error())
				return nil, err
			}
			auth := configYml.GetString("Redis.Password") //通过配置项设置redis密码
			if len(auth) >= 1 {
				if _, err := conn.Do("AUTH", auth); err != nil {
					_ = conn.Close()
					system.ZapLog.Error(consts.ErrorsRedisAuthFail + err.Error())
				}
			}
			_, _ = conn.Do("select", configYml.GetInt("Redis.IndexDb"))
			return conn, err
		},
	}

	conn := redisPool.Get()
	defer conn.Close()

	return redisPool
}

func Exists(key string) bool {
	conn := initRedisClientPool().Get()
	defer conn.Close()
	if conn == nil {
		panic("redis init failed")
	}

	n, err := redis.Int(conn.Do("EXISTS", key))

	if err != nil {
		return false
	}
	return n > 0
}

func Set(key string, value string, refresh bool, ttl int) bool {
	conn := initRedisClientPool().Get()
	defer conn.Close()

	exists := Exists(key)

	if exists && !refresh {
		return false
	}

	if refresh && exists {
		var err error
		if ttl != -1 {
			_, err = redis.String(conn.Do("set", key, value, "ex", ttl))
		} else {
			_, err = redis.String(conn.Do("set", key, value))
		}

		if err != nil {
			return false
		}

		return true
	}

	if !exists {
		var err error

		if ttl != -1 {
			_, err = redis.String(conn.Do("set", key, value, "ex", ttl))
		} else {
			_, err = redis.String(conn.Do("set", key, value))
		}

		if err != nil {
			return false
		}

		return true
	}

	return false
}

func Get(key string) (bytes []byte, err error) {

	conn := initRedisClientPool().Get()
	defer conn.Close()

	val, err := conn.Do("get", key)
	if err != nil {
		return nil, err
	}
	if value, ok := val.([]byte); ok {
		return value, nil
	}
	return []byte{}, nil
}

// 获取 过期时间
func TTL(key string) (time int, err error) {
	conn := initRedisClientPool().Get()

	defer conn.Close()

	time, err = redis.Int(conn.Do("ttl", key))
	return
}

func Ping() {

	conn := initRedisClientPool().Get()
	defer conn.Close()

	r, err := redis.String(conn.Do("PING", "test"))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[debug] redis connect failed %s", err.Error())
	}

	if r != "test" {
		_, _ = fmt.Fprintf(os.Stderr, "[debug] redis connect failed")
		os.Exit(-1)
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "[debug] redis connect success")
	}
}

func LPush(key string, value string) (info int64, err error) {
	conn := initRedisClientPool().Get()
	defer conn.Close()
	info, err = redis.Int64(conn.Do("LPUSH", key, value))
	return
}

func BRPop(key string) (bates []string, err error) {
	// BRPOP key1 [key2 ] timeout 0为 永久不堵塞
	conn := initRedisClientPool().Get()
	defer conn.Close()
	val, err := conn.Do("BRPOP", key, 0)
	// val  实际上是  []interface{} 里面 包着两个 interface{}
	if err != nil {
		return nil, err
	}

	if value, ok := val.([]interface{}); ok {
		for _, values := range value {
			if vals, ok := values.([]byte); ok {
				bates = append(bates, string(vals))
			}
		}
	}

	return
}

func RPop(key string) (bates []byte, err error) {
	// RPop key1
	conn := initRedisClientPool().Get()
	defer conn.Close()
	val, err := conn.Do("RPOP", key)

	if err != nil {
		return nil, err
	}
	if value, ok := val.([]byte); ok {
		return value, nil
	}
	return []byte{}, nil
}

// Del 删除redis中 单个的key
func Del(key string) (time int, err error) {
	conn := initRedisClientPool().Get()
	defer conn.Close()

	time, err = redis.Int(conn.Do("del", key))
	return
}
