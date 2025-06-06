@taskkill /f /im main.exe >nul 2>&1

@REM go install github.com/swaggo/swag/cmd/swag@latest
@REM go mod tidy

swag init

go env -w GOOS=windows
go run main.go