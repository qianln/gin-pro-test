package consts

const (
	// 进程被结束
	ProcessKilled string = "收到信号，进程被结束"

	Param                        string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误, "
	GinSetTrustProxyError   string = "Gin 设置信任代理服务器出错"

	JwtTokenOK            int    = 200100                      //token有效
	JwtTokenInvalid       int    = -400100                     //无效的token
	JwtTokenExpired       int    = -400101                     //过期的token
	JwtTokenFormatErrCode int    = -400102                     //提交的 token 格式错误
	JwtTokenFormatErrMsg  string = "提交的 token 格式错误"            //提交的 token 格式错误
	JwtTokenMustValid     string = "token为必填项,请在请求header部分提交!" //提交的 token 格式错误

	// StartTimeStamp SnowFlake 雪花算法
	StartTimeStamp = int64(1483228800000) //开始时间截 (2017-01-01)
	MachineIdBits  = uint(10)             //机器id所占的位数
	SequenceBits   = uint(12)             //序列所占的位数
	//MachineIdMax   = int64(-1 ^ (-1 << MachineIdBits)) //支持的最大机器id数量
	SequenceMask   = int64(-1 ^ (-1 << SequenceBits)) //
	MachineIdShift = SequenceBits                     //机器id左移位数
	TimestampShift = SequenceBits + MachineIdBits     //时间戳左移位数

	CurdStatusOkCode         int    = 200
	CurdStatusOkMsg          string = "Success"
	CurdCreatFailCode        int    = -400200
	CurdCreatFailMsg         string = "新增失败"
	CurdUpdateFailCode       int    = -400201
	CurdUpdateFailMsg        string = "更新失败"
	CurdDeleteFailCode       int    = -400202
	CurdDeleteFailMsg        string = "删除失败"
	CurdSelectFailCode       int    = -400203
	CurdSelectFailMsg        string = "查询无数据"
	CurdRegisterFailCode     int    = -400204
	CurdRegisterFailMsg      string = "注册失败"
	CurdRegisterUserExistMsg string = "用户已存在"
	CurdLoginFailCode        int    = -400205
	CurdLoginFailMsg         string = "登录失败,请检查用户名或者密码"
	CurdRefreshTokenFailCode int    = -400206
	CurdRefreshTokenFailMsg  string = "刷新Token失败"

	HttpStatusOkCode int    = 200
	HttpStatusOkMsg  string = "Success"

	//文件上传
	FilesUploadFailCode            int    = -400250
	FilesUploadFailMsg             string = "文件上传失败, 获取上传文件发生错误!"
	FilesUploadMoreThanMaxSizeCode int    = -400251
	FilesUploadMoreThanMaxSizeMsg  string = "长传文件超过系统设定的最大值,系统允许的最大值："
	FilesUploadMimeTypeFailCode    int    = -400252
	FilesUploadMimeTypeFailMsg     string = "文件mime类型不允许"

	//websocket
	WsServerNotStartCode int    = -400300
	WsServerNotStartMsg  string = "websocket 服务没有开启，请在配置文件开启，相关路径：config/config.yml"
	WsOpenFailCode       int    = -400301
	WsOpenFailMsg        string = "websocket open阶段初始化基本参数失败"

	//验证码
	CaptchaGetParamsInvalidMsg    string = "获取验证码：提交的验证码参数无效,请检查验证码ID以及文件名后缀是否完整"
	CaptchaGetParamsInvalidCode   int    = -400350
	CaptchaCheckParamsInvalidMsg  string = "校验验证码：提交的参数无效，请检查 【验证码ID、验证码值】 提交时的键名是否与配置项一致"
	CaptchaCheckParamsInvalidCode int    = -400351
	CaptchaCheckOkMsg             string = "验证码校验通过"
	CaptchaCheckFailCode          int    = -400355
	CaptchaCheckFailMsg           string = "验证码校验失败"
)

const (
	//系统部分
	ErrorsContainerKeyAlreadyExists string = "该键已经注册在容器中了"
	ErrorsPublicNotExists           string = "public 目录不存在"
	ErrorsConfigYamlNotExists       string = "config.yml 配置文件不存在"
	ErrorsConfigGormNotExists       string = "gorm_v2.yml 配置文件不存在"
	ErrorsStorageLogsNotExists      string = "storage/logs 目录不存在"
	ErrorsConfigInitFail            string = "初始化配置文件发生错误"
	ErrorsSoftLinkCreateFail        string = "自动创建软连接失败,请以管理员身份运行客户端(开发环境为goland等，生产环境检查命令执行者权限), " +
		"最后一个可能：如果您是360用户，请退出360相关软件，才能保证go语言创建软连接函数： os.Symlink() 正常运行"
	ErrorsSoftLinkDeleteFail string = "删除软软连接失败"

	ErrorsFuncEventAlreadyExists   string = "注册函数类事件失败，键名已经被注册"
	ErrorsFuncEventNotRegister     string = "没有找到键名对应的函数"
	ErrorsFuncEventNotCall         string = "注册的函数无法正确执行"
	ErrorsBasePath                 string = "初始化项目根目录失败"
	ErrorsTokenBaseInfo            string = "token最基本的格式错误,请提供一个有效的token!"
	ErrorsNoAuthorization          string = "token鉴权未通过，请通过token授权接口重新获取token,"
	ErrorsRefreshTokenFail         string = "token不符合刷新条件,请通过登陆接口重新获取token!"
	ErrorsParseTokenFail           string = "解析token失败"
	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"
	ErrorsCasbinNoAuthorization    string = "Casbin 鉴权未通过，请在后台检查 casbin 设置参数"
	ErrorsGormNotInitGlobalPointer string = "%s 数据库全局变量指针没有初始化，请在配置文件 config/gorm_v2.yml 设置 Gormv2.%s.IsInitGlobalGormMysql = 1, 并且保证数据库配置正确 \n"

	// 数据库部分
	ErrorsDbDriverNotExists        string = "数据库驱动类型不存在,目前支持的数据库类型：mysql、sqlserver、postgresql，您提交数据库类型："
	ErrorsDialectorDbInitFail      string = "gorm dialector 初始化失败,dbType:"
	ErrorsGormDBCreateParamsNotPtr string = "gorm Create 函数的参数必须是一个指针"
	ErrorsGormDBUpdateParamsNotPtr string = "gorm 的 Update、Save 函数的参数必须是一个指针(GinSkeleton ≥ v1.5.29 版本新增验证，为了完美支持 gorm 的所有回调函数,请在参数前面添加 & )"
	//redis部分
	ErrorsRedisInitConnFail string = "初始化redis连接池失败"
	ErrorsRedisAuthFail     string = "Redis Auth 鉴权失败，密码错误"
	ErrorsRedisGetConnFail  string = "Redis 从连接池获取一个连接失败，超过最大重试次数"
	// 表单参数验证器未通过时的错误
	ErrorsValidatorNotExists      string = "不存在的验证器"
	ErrorsValidatorTransInitFail  string = "validator的翻译器初始化错误"
	ErrorNotAllParamsIsBlank      string = "该接口不允许所有参数都为空,请按照接口要求提交必填参数"
	ErrorsValidatorBindParamsFail string = "验证器绑定参数失败"

	//token部分
	ErrorsTokenInvalid      string = "无效的token"
	ErrorsTokenNotActiveYet string = "token 尚未激活"
	ErrorsTokenMalFormed    string = "token 格式不正确"

	//snowflake
	ErrorsSnowflakeGetIdFail string = "获取snowflake唯一ID过程发生错误"

	//文件上传
	ErrorsFilesUploadOpenFail string = "打开文件失败，详情："
	ErrorsFilesUploadReadFail string = "读取文件32字节失败，详情："
)
