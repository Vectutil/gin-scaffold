package utils

import (
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"net"
	"strings"
)

func RunInfo() {
	ip := GetLocalIP()
	port := config.Cfg.System.Port
	goUrl := fmt.Sprintf("http://%s:%s", ip, port)
	swaggerURL := fmt.Sprintf("http://%s:%s/swagger/index.html#/", ip, port)
	logger.Logger.Info("项目启动成功")
	logger.Logger.Info(fmt.Sprintf("项目启动成功，go  地 址: %s", goUrl))
	logger.Logger.Info(fmt.Sprintf("项目启动成功，swag地址为: %s", swaggerURL))

}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip
}
