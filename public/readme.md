### 关于 GIN_PRO
>
### 功能介绍

>

### 目录结构

```
├── app -- 核心代码
│   ├── console --任务调度
│   │   ├── commands
│   │   │   └── test_log.go
│   │   └── kernel.go
│   ├── core    -- 核心结构  
│   │   ├── container
│   │   │   └── container.go
│   │   └── system
│   │       └── system.go
│   ├── global      -- 自定义常量
│   │   └── consts
│   │       └── consts.go
│   ├── http    -- http请求
│   │   ├── controllers
│   │   │   ├── base_controller.go
│   │   │   └── open
│   │   │       └── demo_controller.go
│   │   ├── kernel.go
│   │   ├── middleware
│   │   │   ├── authorization
│   │   │   │   └── auth.go
│   │   │   └── cors
│   │   │       └── cors.go
│   │   └── validator
│   │       ├── common_data
│   │       │   └── common_data.go
│   │       └── open
│   │           └── demo
│   │               └── index_validator.go
│   ├── modules     -- 公用服务层
│   │   ├── models
│   │   │   ├── base_model.go
│   │   │   └── users_model.go
│   │   └── services
│   │       └── base_services.go
│   └── utils   -- 工具模块
│       ├── helps
│       │   └── helps.go
│       ├── redis
│       │   └── client.go
│       └── response
│           └── response.go
├── bootstrap   -- 项目启动初始化
│   └── init.go 
├── config  -- 配置文件信息
│   └── config.yml
├── database    -- 数据库相关
│   └── readme.md
├── go.mod
├── go.sum
├── library -- 外部依赖封装
│   ├── config
│   │   ├── config.go
│   │   └── iconfig
│   │       └── iconfig.go
│   ├── mysql_gorm
│   │   ├── client.go
│   │   ├── custom_log.go
│   │   └── hook.go
│   ├── release_router
│   │   └── release_router.go
│   ├── snow_flake
│   │   └── snow_flake.go
│   ├── validator_translation
│   │   └── validator_transiation.go
│   └── zap_log
│       ├── zap_log.go
│       └── zap_log_hook
│           └── zap_log_hooks.go
├── main.go
├── public    -- 静态文件 
│   ├── css
│   │   └── layui.css
│   ├── font
│   │   ├── iconfont.eot
│   │   ├── iconfont.svg
│   │   ├── iconfont.ttf
│   │   ├── iconfont.woff
│   │   └── iconfont.woff2
│   ├── js
│   │   ├── jquery-1.11.3.min.js
│   │   └── layui.js
│   ├── readme.md
│   └── storage -> /Users/qianln/go/src/gin-pro/storage/app
├── routers     -- 路由模块 
│   ├── app.go
│   └── open
│       └── demo.go
├── storage     --日志文件及资源文件
│   ├── app
│   │   └── imgs
│   │       └── img.png
│   └── logs
├── template    -- 模板文件
│   └── index.html

```

