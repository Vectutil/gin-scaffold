# 使用 Unix 风格路径
export PATH := C:/Go/bin:$(PATH)

.PHONY: run

run:
	@taskkill /f /im main.exe >nul 2>&1
	@swag init
	@echo "Using go from: $(shell which go)"
	@go run main.go
