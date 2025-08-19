package windows_send

import (
	"fmt"
	"gin-scaffold/pkg/logger"
)

type SendMsgWindows struct {
	WindowsName string
	Msg         string
}

var WindowsSendChan = make(chan SendMsgWindows, 10000)

func SendMsgToWindows(windowsName string, msg string) {

	windows := SendMsgWindows{
		WindowsName: windowsName,
		Msg:         msg,
	}

	WindowsSendChan <- windows
	logger.Logger.Info(fmt.Sprintf("发送消息: %s", msg))

}

func ConsumeMsgWindows() {
	for windows := range WindowsSendChan {
		windowsName := windows.WindowsName
		msg := windows.Msg

		hwnd, found := FindWindow(windowsName)
		if !found {
			logger.Logger.Info(fmt.Sprintf("未找到标题包含 '%s' 的微信窗口", windowsName))
			return
		}

		success := SendWechatMessage(hwnd, msg)
		if success {
			logger.Logger.Info(fmt.Sprintf("成功向 '%s' 发送消息: %s", windowsName, msg))
		} else {
			logger.Logger.Info(fmt.Sprintf("发送消息失败"))
		}
	}
}
