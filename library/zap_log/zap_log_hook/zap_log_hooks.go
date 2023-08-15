package sys_log_hook

import (
	"go.uber.org/zap/zapcore"
)

func ZapLogHandler(entry zapcore.Entry) error {
	/**
		拦截住每一条日志可以做相关处理 推送到阿里云日志管理面板、ElasticSearch 日志库等

		paramEntry 参数介绍
		Level      日志等级
		Time       当前时间
		LoggerName  日志名称
		Message    日志内容
		Caller     各个文件调用路径
		Stack      代码调用栈

		fmt.Printf("%#+v\n", paramEntry)
	*/
	go func(paramEntry zapcore.Entry) {

	}(entry)
	return nil
}
