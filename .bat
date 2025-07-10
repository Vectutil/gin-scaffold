swag init

go env -w GOOS=windows
go build -ldflags "-s -w" main.go