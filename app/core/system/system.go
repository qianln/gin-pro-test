package system

import (
	"gin-pro/app/global/consts"
	"gin-pro/library/config/iconfig"
	"gin-pro/library/snow_flake"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var (
	BasePath              string      // 定义项目的根目录
	ConfigKeyPrefix       = "Config_" // 配置文件键值缓存时，键的前缀
	ParamPrefix           = "Request_Param_"
	DateFormat            = "2006-01-02 15:04:05"     // 设置全局日期时间格式
	DateFormatMilliSecond = "2006-01-02 15:04:05.000" // 设置全局日期时间格式

	ZapLog       *zap.Logger           // 全局日志指针
	Config       iconfig.IConfig       // 加载全局配置文件
	DbMysql      *gorm.DB              // 全局gorm的客户端连接
	SnowFlake    *snow_flake.Snowflake // 雪花算法全局变量
	CobraCommand *cobra.Command        // cli 模式
)

func init() {
	if curPath, err := os.Getwd(); err == nil {
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	} else {
		log.Fatal(consts.ErrorsBasePath)
	}
}
