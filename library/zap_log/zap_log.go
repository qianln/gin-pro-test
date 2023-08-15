package zap_log

import (
	"fmt"
	"gin-pro/app/core/system"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func NewZapLog(entry func(zapcore.Entry) error) *zap.Logger {

	if system.Config.GetBool("AppDebug") == true {
		if logger, err := zap.NewDevelopment(zap.Hooks(entry)); err == nil {
			return logger
		} else {
			log.Fatal("创建zap日志包失败，详情：" + err.Error())
		}
	}

	encoderConfig := zap.NewProductionEncoderConfig()

	// 设置日志存储时间粒度
	timePrecision := system.Config.GetString("Logs.TimePrecision")
	var recordTimeFormat string

	switch timePrecision {
	case "second":
		recordTimeFormat = system.DateFormat
	case "millisecond":
		recordTimeFormat = system.DateFormatMilliSecond
	default:
		recordTimeFormat = system.DateFormat

	}
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "created_at"

	// 设置日志存储格式
	textFormat := system.Config.GetString("Logs.TextFormat")
	var encoder zapcore.Encoder
	switch textFormat {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	//写入器
	timeFormat := time.Now().Format("20060102")
	logName := fmt.Sprintf("/gin-pro-%s.log", timeFormat)
	fileName := system.BasePath + system.Config.GetString("Logs.GinProLogPath") + logName
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                                // 日志文件的位置
		MaxSize:    system.Config.GetInt("Logs.MaxSize"),    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: system.Config.GetInt("Logs.MaxBackups"), // 保留旧文件的最大个数
		MaxAge:     system.Config.GetInt("Logs.MaxAge"),     // 保留旧文件的最大天数
		Compress:   system.Config.GetBool("Logs.Compress"),  // 是否压缩/归档旧文件
	}
	writer := zapcore.AddSync(lumberJackLogger)
	// 开始初始化zap日志核心参数，
	// 参数一：编码器
	// 参数二：写入器
	// 参数三：参数级别，debug级别支持后续调用的所有函数写日志，如果是 fatal 高级别，则级别>=fatal 才可以写日志
	zapCore := zapcore.NewCore(encoder, writer, zap.InfoLevel)
	return zap.New(zapCore, zap.AddCaller(), zap.Hooks(entry), zap.AddStacktrace(zap.WarnLevel))
}
