### 关于 GIN_PRO
>
### 功能介绍

>

### 目录结构

```
├── GinPro
├── app
│   ├── console
│   │   ├── commands
│   │   │   └── test_log.go
│   │   └── kernel.go
│   ├── core
│   │   └── container
│   │       └── container.go
│   ├── global
│   │   ├── consts
│   │   │   └── consts.go
│   │   ├── errors
│   │   │   └── errors.go
│   │   └── variable
│   │       └── variable.go
│   ├── http
│   │   ├── controllers
│   │   │   ├── base_controller.go
│   │   │   ├── open
│   │   │   │   └── demo_controller.go
│   │   │   └── web
│   │   │       └── users_controller.go
│   │   ├── middleware
│   │   │   ├── authorization
│   │   │   │   └── auth.go
│   │   │   └── cors
│   │   │       └── cors.go
│   │   └── validator
│   │       ├── common
│   │       │   ├── base_data
│   │       │   │   └── base_data.go
│   │       │   └── register_validator
│   │       │       └── wap_register_validator.go
│   │       ├── core
│   │       │   ├── data_transfer
│   │       │   │   └── data_transfer.go
│   │       │   ├── http_request
│   │       │   │   └── http_request.go
│   │       │   └── validator_interface
│   │       │       └── validator_interface.go
│   │       └── web
│   │           └── users
│   │               ├── login_validator.go
│   │               └── register_validator.go
│   ├── models
│   │   ├── base_model.go
│   │   └── users_model.go
│   ├── services
│   │   ├── base_services.go
│   │   └── web
│   │       └── users_services.go
│   └── utils
│       ├── helps
│       │   └── helps.go
│       ├── redis
│       │   └── client.go
│       └── response
│           └── response.go
├── bootstrap
│   └── init.go
├── config
│   └── config.yml
├── database
│   └── readme.md
├── go.mod
├── go.sum
├── library
│   ├── config
│   │   ├── config.go
│   │   └── config_interface
│   │       └── config_interface.go
│   ├── my_jwt
│   │   └── my_jwt.go
│   ├── mysql_gorm
│   │   ├── client.go
│   │   ├── custom_log.go
│   │   └── hook.go
│   ├── release_router
│   │   └── release_router.go
│   ├── snow_flake
│   │   └── snow_flake.go
│   ├── validator_translation
│   │   └── validator_transiation.go
│   └── zap_log
│       ├── zap_log.go
│       └── zap_log_hook
│           └── zap_log_hooks.go
├── main.go
├── public
│   └── readme.md
├── routers
│   ├── api
│   ├── app.go
│   ├── open
│   │   └── demo.go
│   └── web
│       └── users.go
├── storage
│   ├── app
│   │   └── imgs
│   │       └── img.png
│   └── logs
│       ├── GinPro-20230407.log
│       └── GinPro-20230412.log
├── template
│   └── index.tmpl
└── tmp
    └── tmp_bin


```

