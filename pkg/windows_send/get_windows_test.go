package windows_send

import (
	"fmt"
	"testing"
)

func TestGetWindows(t *testing.T) {
	// 目标聊天窗口名称
	// 查找微信聊天窗口（标题包含联系人名称）
	contactName := "三傻大闹好莱坞" // 替换为实际联系人名称
	hwnd, found := FindWindow(contactName)
	if !found {
		fmt.Printf("未找到标题包含 '%s' 的微信窗口\n", contactName)
		return
	}

	// 发送消息
	message := "这是一条通过程序发送的消息"
	success := SendWechatMessage(hwnd, message)
	if success {
		fmt.Printf("成功向 '%s' 发送消息: %s\n", contactName, message)
	} else {
		fmt.Println("发送消息失败")
	}
}
