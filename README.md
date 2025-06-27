# gin-scaffold
    gin 脚手架 旨在开箱即用 分层清晰 按需使用-- (逐步完善)
    更多说明在语雀文档：https://www.yuque.com/juz652/ow4qcx?# 《程序项目》

    中间件：
        gorm        持久层
        zap         日志      (github.com/robfig/cron/v3)
        job         定时任务
        redis       缓存      (github.com/redis/go-redis/v9)
        RabbitMQ    消息队列

        qny         七牛云
        swag        swagger文档
        

    等 .... 完善中

# 权限管理
    
```
    登录管理:
    将登录信息存进redis中
    生成2小时有效的accessToken 防止token被盗用
```
    使用中间件判断权限、状态


    权限模型设计要素
    用户（users）
    角色（roles）
    菜单/页面模块（menus） — 包含按钮操作
    权限（permissions） — 功能点级别（如按钮）
    角色-权限映射（role_permissions）
    用户-角色映射（user_roles）
    数据权限（data_scope） — 控制某个用户或角色可以访问哪些数据（如部门、区域等）

# 定时任务
    1. 引入
        go get github.com/robfig/cron/v3
    2. config文件编写
        job:
            jobStatus:
            examplejob: true             // 是否要开启这个定时任务
            
            jobCron:
            examplejob: "*/10 * * * * *" // 每10秒执行一次
    3. 前往./internal/app/job目录下编写定时任务
        

# 运行方式
    window用户：
        第一次运行需要在 .bat 中增加命令代码
            ```
            go install github.com/swaggo/swag/cmd/swag@latest
            go mod tidy
            ```
    控制台执行 .bat
    
    go env -w GOOS=linux
    go env -w GOOS=windows
# 结构
    1. 项目结构

```
gin-scaffold/
├── main.go               # 启动入口
├── internal/             # 内部实现模块（非导出）
│   ├── app/              # 核心业务逻辑划分模块
│   │   ├── consts/       # 常量
│   │   ├── dao/          # 数据访问层（数据库操作）
│   │   ├── handler/      # 控制器层（HTTP 请求处理）
│   │   ├── job/          # 定时任务或异步任务
│   │   ├── logic/        # 核心业务逻辑
│   │   ├── model/        # 数据模型定义
│   │   ├── response/     # gin 成功或错误响应
│   │   └── types/        # 响应结构体定义
│   ├── config/           # 配置文件加载与管理
│   ├── middleware/       # Gin 中间件
│   └── router/           # 路由定义
├── pkg/                  # 可复用工具库（非业务相关，独立通用）
│   ├── http_call/        # HTTP 通用封装
│   ├── logger/           # 日志封装
│   ├── mysql/            # 数据库连接/初始化
│   ├── qny/              # qny 封装
│   ├── rabbitmq/         # rabbitmq 封装 （延迟队列丨普通队列）
│   ├── redis/            # redis 连接/使用
│   └── utils             # 工具包
├── cache/                # 缓存文件(日志等)
├── docs/                 # swagger文档
├── .bat                  # 启动脚本(window)
├── makefile              # 启动脚本(linux)
├── config.yaml           # 配置文件
├── main.go               # 启动入口
├── go.mod                # 依赖管理
└── README.md  
```

# 日志 zap
