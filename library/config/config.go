package config

import (
	"gin-pro/app/core/container"
	"gin-pro/app/core/system"
	"gin-pro/app/global/consts"
	"gin-pro/library/config/iconfig"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var lastChangeTime time.Time
var containers = container.NewContainers()

func init() {
	lastChangeTime = time.Now()
}

// 判断相关键是否已经缓存
func (y *Config) keyIsCache(keyName string) bool {
	_, exists := containers.Exists(system.ConfigKeyPrefix + keyName)
	return exists
}

// 对键值进行缓存
func (y *Config) cache(keyName string, value interface{}) bool {
	// 避免瞬间缓存键、值时，程序提示键名已经被注册的日志输出
	y.mu.Lock()
	defer y.mu.Unlock()
	if _, exists := containers.Exists(system.ConfigKeyPrefix + keyName); exists {
		return true
	}
	return containers.Set(system.ConfigKeyPrefix+keyName, value)
}

// 通过键获取缓存的值
func (y *Config) getValueFromCache(keyName string) interface{} {
	return containers.Get(system.ConfigKeyPrefix + keyName)
}

// 清空已经缓存的配置项信息
func (y *Config) clearCache() {
	containers.FuzzyDelete(system.ConfigKeyPrefix)
}

func NewConfig() iconfig.IConfig {
	config := viper.New()

	config.AddConfigPath(system.BasePath + "/config")
	config.SetConfigName("config")
	config.SetConfigType("yml")

	if err := config.ReadInConfig(); err != nil {
		log.Fatal(consts.ErrorsConfigInitFail + err.Error())
	}

	return &Config{
		viper: config,
		mu:    new(sync.Mutex),
	}
}

type Config struct {
	viper *viper.Viper
	mu    *sync.Mutex
}

func (y *Config) ConfigFileChangeListen() {
	y.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			if changeEvent.Op.String() == "WRITE" {
				y.clearCache()
				lastChangeTime = time.Now()
			}
		}
	})
	y.viper.WatchConfig()
}

func (y *Config) Get(keyName string) any {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName)
	} else {
		value := y.viper.Get(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetString(keyName string) string {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(string)
	} else {
		value := y.viper.GetString(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetBool(keyName string) bool {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(bool)
	} else {
		value := y.viper.GetBool(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetInt(keyName string) int {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(int)
	} else {
		value := y.viper.GetInt(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetInt32(keyName string) int32 {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(int32)
	} else {
		value := y.viper.GetInt32(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetInt64(keyName string) int64 {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(int64)
	} else {
		value := y.viper.GetInt64(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetFloat64(keyName string) float64 {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(float64)
	} else {
		value := y.viper.GetFloat64(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetDuration(keyName string) time.Duration {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(time.Duration)
	} else {
		value := y.viper.GetDuration(keyName)
		y.cache(keyName, value)
		return value
	}
}

func (y *Config) GetStringSlice(keyName string) []string {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).([]string)
	} else {
		value := y.viper.GetStringSlice(keyName)
		y.cache(keyName, value)
		return value
	}
}
