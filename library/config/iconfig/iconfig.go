package iconfig

import "time"

// IConfig 主要解决 循环引用问题
type IConfig interface {

	//ConfigFileChangeListen 监听文件变化
	ConfigFileChangeListen()

	// Clone 允许 clone 一个相同功能的结构体
	Clone(fileName string) IConfig

	// Get 一个原始值
	Get(keyName string) interface{}

	// GetString 字符串格式返回值
	GetString(keyName string) string

	// GetBool 布尔格式返回值
	GetBool(keyName string) bool

	// GetInt 整数格式返回值
	GetInt(keyName string) int

	// GetInt32 整数格式返回值
	GetInt32(keyName string) int32

	// GetInt64 整数格式返回值
	GetInt64(keyName string) int64

	// GetFloat64 小数格式返回值
	GetFloat64(keyName string) float64

	// GetDuration 时间单位格式返回值
	GetDuration(keyName string) time.Duration

	// GetStringSlice 字符串切片数格式返回值
	GetStringSlice(keyName string) []string
}
