package container

import (
	"gin-pro/app/core/system"
	"gin-pro/app/global/consts"
	"log"
	"strings"
	"sync"
)

var sMap sync.Map

func NewContainers() *Containers {
	return &Containers{}
}

type Containers struct {
}

// Set  1.以键值对的形式将代码注册到容器
func (c *Containers) Set(key string, value interface{}) (res bool) {

	if _, exists := c.Exists(key); exists == false {
		sMap.Store(key, value)
		res = true
	} else {

		if system.ZapLog == nil {
			log.Fatal(consts.ErrorsContainerKeyAlreadyExists + ",请解决键名重复问题,相关键：" + key)
		} else {
			system.ZapLog.Warn(consts.ErrorsContainerKeyAlreadyExists + ", 相关键：" + key)
		}
	}
	return
}

// Delete  2.删除
func (c *Containers) Delete(key string) {
	sMap.Delete(key)
}

// Get 3.传递键，从容器获取值
func (c *Containers) Get(key string) interface{} {
	if value, exists := c.Exists(key); exists {
		return value
	}
	return nil
}

// Exists 4. 判断键是否被注册
func (c *Containers) Exists(key string) (interface{}, bool) {
	return sMap.Load(key)
}

// FuzzyDelete 按照键的前缀模糊删除容器中注册的内容
func (c *Containers) FuzzyDelete(keyPre string) {
	sMap.Range(func(key, value interface{}) bool {
		if keyname, ok := key.(string); ok {
			if strings.HasPrefix(keyname, keyPre) {
				sMap.Delete(keyname)
			}
		}
		return true
	})
}
