package bootstrap

import (
	"gin-pro/app/console"
	"gin-pro/app/core/system"
	"gin-pro/app/global/consts"
	"gin-pro/library/cobra_command"
	"gin-pro/library/config"
	"gin-pro/library/mysql_gorm"
	"gin-pro/library/snow_flake"
	"gin-pro/library/validator_translation"
	"gin-pro/library/zap_log"
	"gin-pro/library/zap_log/zap_log_hook"
	"log"
	"os"
)

func init() {

	// 1. 检查必要文件是否存在
	checkRequiredFolders()

	// 2.读取配置文件到容器内 监听配置文件变化
	system.Config = config.NewConfig()

	// 3.初始化全局日志句柄，并载入日志钩子处理函数
	system.ZapLog = zap_log.NewZapLog(sys_log_hook.ZapLogHandler)

	// 4.初始化 mysql
	dbMysql, err := mysql_gorm.GetOneMysqlClient()
	if err != nil {
		log.Fatal(consts.ErrorsGormInitFail + err.Error())
		return
	}
	system.DbMysql = dbMysql

	// 5.雪花算法全局变量
	system.SnowFlake = snow_flake.NewSnowFlake()

	// 6.全局注册 [validator 错误翻译器,zh 代表中文，en 代表英语]
	if err := validator_translation.InitTrans("zh"); err != nil {
		log.Fatal(consts.ErrorsValidatorTransInitFail + err.Error())
	}

	// 7.注册计划任务
	console.Schedule()

	// 注册全局命令行参数
	system.CobraCommand = cobra_command.NewCobraCommand()

}

func checkRequiredFolders() {
	if _, err := os.Stat(system.BasePath + "/config/config.yml"); err != nil {
		log.Fatal(consts.ErrorsConfigYamlNotExists + err.Error())
	}

	if _, err := os.Stat(system.BasePath + "/public/"); err != nil {
		log.Fatal(consts.ErrorsPublicNotExists + err.Error())
	}

	if _, err := os.Stat(system.BasePath + "/storage/logs/"); err != nil {
		log.Fatal(consts.ErrorsStorageLogsNotExists + err.Error())
	}

	if _, err := os.Stat(system.BasePath + "/public/storage"); err == nil {
		if err = os.RemoveAll(system.BasePath + "/public/storage"); err != nil {
			log.Fatal(consts.ErrorsSoftLinkDeleteFail + err.Error())
		}
	}

	if err := os.Symlink(system.BasePath+"/storage/app", system.BasePath+"/public/storage"); err != nil {
		log.Fatal(consts.ErrorsSoftLinkCreateFail + err.Error())
	}
}
