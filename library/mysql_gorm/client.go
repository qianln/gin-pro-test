package mysql_gorm

import (
	"fmt"
	"gin-pro/app/core/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func GetOneMysqlClient() (*gorm.DB, error) {

	// 获取连接句柄
	dbDial := mysql.Open(getDsnByConfig())

	gormDb, err := gorm.Open(dbDial, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 redefineLog(),
	})
	if err != nil {
		return nil, err
	}

	_ = gormDb.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", MaskNotDataError)
	_ = gormDb.Callback().Create().Before("gorm:before_create").Register("CreateBeforeHook", CreateBeforeHook)
	_ = gormDb.Callback().Update().Before("gorm:before_update").Register("UpdateBeforeHook", UpdateBeforeHook)

	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(system.Config.GetDuration("Mysql.SetConnMaxLifetime") * time.Second)
		rawDb.SetMaxIdleConns(system.Config.GetInt("Mysql.SetMaxIdleConns"))
		rawDb.SetMaxOpenConns(system.Config.GetInt("Mysql.SetMaxOpenConns"))
		return gormDb, nil
	}

}

func getDsnByConfig() string {
	Host := system.Config.GetString("Mysql.Host")
	Port := system.Config.GetString("Mysql.Port")
	DataBase := system.Config.GetString("Mysql.DataBase")
	User := system.Config.GetString("Mysql.User")
	Pass := system.Config.GetString("Mysql.Pass")
	Charset := system.Config.GetString("Mysql.Charset")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=false&loc=Local&parseTime=true", User, Pass, Host, Port, DataBase, Charset)
}

// 创建自定义日志模块
func redefineLog() gormLogger.Interface {
	return NewCustomLog(
		SetInfoStrFormat("[info] %s\n"),
		SetWarnStrFormat("[warn] %s\n"),
		SetErrStrFormat("[error] %s\n"),
		SetTraceStrFormat("[traceStr] %s [%.3fms] [rows:%v] %s\n"),
		SetTracWarnStrFormat("[traceWarn] %s %s [%.3fms] [rows:%v] %s\n"),
		SetTracErrStrFormat("[traceErr] %s %s [%.3fms] [rows:%v] %s\n"),
	)
}
