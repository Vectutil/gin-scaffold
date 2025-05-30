REM 杀掉之前运行的 main.exe（忽略错误）
@taskkill /f /im main.exe >nul 2>&1

REM 生成 Swagger 文档（确保已安装 swag）
swag init

REM 显示当前使用的 go 路径
where go

REM 启动程序
go run main.go