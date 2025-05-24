# gin-scaffold
gin 脚手架 旨在开箱即用 分层清晰 按需使用-- (逐步完善)
    
    中间件：
        gorm  持久层
        zap   日志      (github.com/robfig/cron/v3)
        job   定时任务
        redis 缓存      (github.com/redis/go-redis/v9)

        qny   七牛云
        

    等 .... 完善中

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
        

# 结构
    1. 项目结构

```
gin-scaffold/
├── main.go               # 启动入口
├── internal/             # 内部实现模块（非导出）
│   ├── app/              # 核心业务逻辑划分模块
│   │   ├── handler/      # 控制器层（HTTP 请求处理）
│   │   ├── logic/        # 核心业务逻辑
│   │   ├── dao/          # 数据访问层（数据库操作）
│   │   ├── model/        # 数据模型定义
│   │   └── job/          # 定时任务或异步任务
│   ├── config/           # 配置文件加载与管理
│   ├── middleware/       # Gin 中间件
│   └── router/           # 路由定义
├── pkg/                  # 可复用工具库（非业务相关，独立通用）
│   ├── http_call/        # HTTP 通用封装
│   ├── logger/           # 日志封装
│   ├── qny/              # qny 封装
│   ├── redis/            # redis 连接/使用
│   └── mysql/            # 数据库连接/初始化
├── cache/                # 缓存文件(日志等)
├── go.mod
├── go.sum
└── README.md
```

# 日志 zap
    