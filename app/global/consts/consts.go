package consts

const (
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误, "

	StartTimeStamp = int64(1483228800000)             //开始时间截 (2017-01-01)
	MachineIdBits  = uint(10)                         //机器id所占的位数
	SequenceBits   = uint(12)                         //序列所占的位数
	SequenceMask   = int64(-1 ^ (-1 << SequenceBits)) //
	MachineIdShift = SequenceBits                     //机器id左移位数
	TimestampShift = SequenceBits + MachineIdBits     //时间戳左移位数

	HttpStatusOkCode int    = 200
	HttpStatusOkMsg  string = "Success"
)

const (
	ErrorsContainerKeyAlreadyExists string = "该键已经注册在容器中了"
	ErrorsPublicNotExists           string = "public 目录不存在"
	ErrorsConfigYamlNotExists       string = "config.yml 配置文件不存在"
	ErrorsStorageLogsNotExists      string = "storage/logs 目录不存在"
	ErrorsConfigInitFail            string = "初始化配置文件发生错误"
	ErrorsSoftLinkCreateFail        string = "自动创建软连接失败,请以管理员身份运行客户端(开发环境为goland等，生产环境检查命令执行者权限), "
	ErrorsSoftLinkDeleteFail        string = "删除软连接失败"

	ErrorsBasePath                 string = "初始化项目根目录失败"
	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"
	ErrorsGormNotInitGlobalPointer string = "数据库全局变量指针没有初始化"
	ErrorsGormDBCreateParamsNotPtr string = "gorm Create 函数的参数必须是一个指针"
	ErrorsGormDBUpdateParamsNotPtr string = "gorm 的 Update、Save 函数的参数必须是一个指针"

	ErrorsRedisInitConnFail string = "初始化redis连接池失败"
	ErrorsRedisAuthFail     string = "Redis Auth 鉴权失败，密码错误"
	ErrorsRedisGetConnFail  string = "Redis 从连接池获取一个连接失败，超过最大重试次数"

	ErrorsValidatorTransInitFail string = "validator的翻译器初始化错误"
	ErrorNotAllParamsIsBlank     string = "该接口不允许所有参数都为空,请按照接口要求提交必填参数"

	ErrorsSnowflakeGetIdFail string = "获取snowflake唯一ID过程发生错误"

	ErrorsFilesUploadOpenFail string = "打开文件失败，详情："
	ErrorsFilesUploadReadFail string = "读取文件32字节失败，详情："
)
