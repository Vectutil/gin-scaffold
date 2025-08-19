package windows_send

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	enumWindows         = user32.NewProc("EnumWindows")
	getWindowText       = user32.NewProc("GetWindowTextW")
	setForegroundWindow = user32.NewProc("SetForegroundWindow")
	sendMessageW        = user32.NewProc("SendMessageW")
	openClipboard       = user32.NewProc("OpenClipboard")
	closeClipboard      = user32.NewProc("CloseClipboard")
	emptyClipboard      = user32.NewProc("EmptyClipboard")
	setClipboardData    = user32.NewProc("SetClipboardData")
	globalAlloc         = kernel32.NewProc("GlobalAlloc")
	globalLock          = kernel32.NewProc("GlobalLock")
	globalUnlock        = kernel32.NewProc("GlobalUnlock")
	procKeybdEvent      = user32.NewProc("keybd_event")
)

const (
	WM_SETTEXT     = 0x000C
	WM_KEYDOWN     = 0x0100
	WM_KEYUP       = 0x0101
	VK_RETURN      = 0x0D
	GMEM_MOVEABLE  = 0x0002
	CF_UNICODETEXT = 13

	VK_CONTROL      = 0x11
	VK_V            = 0x56
	KEYEVENTF_KEYUP = 0x0002
)

type WindowInfo struct {
	Hwnd  syscall.Handle
	Title string
}

var targetWindowTitle string
var foundWindows []WindowInfo

// 枚举窗口回调函数
func enumWindowsProc(hwnd syscall.Handle, lParam uintptr) uintptr {
	buf := make([]uint16, 256)
	_, _, _ = getWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)

	title := syscall.UTF16ToString(buf)
	if title != "" && contains(title, targetWindowTitle) {
		foundWindows = append(foundWindows, WindowInfo{Hwnd: hwnd, Title: title})
	}
	return 1
}

// 字符串包含检查
func contains(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// 查找窗口
func FindWindow(title string) (syscall.Handle, bool) {
	targetWindowTitle = title
	foundWindows = []WindowInfo{}
	_, _, _ = enumWindows.Call(syscall.NewCallback(enumWindowsProc), 0)

	if len(foundWindows) > 0 {
		return foundWindows[0].Hwnd, true
	}
	return 0, false
}

// 设置剪贴板文本
func SetClipboardText(text string) bool {
	// 打开剪贴板
	ret, _, _ := openClipboard.Call(0)
	if ret == 0 {
		return false
	}
	defer closeClipboard.Call()

	// 清空剪贴板
	emptyClipboard.Call()

	// 分配内存
	unicodeText := syscall.StringToUTF16(text)
	size := uintptr(len(unicodeText) * 2)
	hMem, _, _ := globalAlloc.Call(GMEM_MOVEABLE, size)
	if hMem == 0 {
		return false
	}

	// 锁定内存并复制数据
	lpMem, _, _ := globalLock.Call(hMem)
	if lpMem == 0 {
		return false
	}
	defer globalUnlock.Call(hMem)

	copy((*[1 << 20]uint16)(unsafe.Pointer(lpMem))[:], unicodeText)

	// 设置剪贴板数据
	ret, _, _ = setClipboardData.Call(CF_UNICODETEXT, hMem)
	return ret != 0
}

// 向窗口发送按键
func SendKey(hwnd syscall.Handle, key uintptr) {
	sendMessageW.Call(uintptr(hwnd), WM_KEYDOWN, key, 0)
	time.Sleep(50 * time.Millisecond)
	sendMessageW.Call(uintptr(hwnd), WM_KEYUP, key, 0)
	time.Sleep(50 * time.Millisecond)
}

// 发送消息到微信窗口
func SendWechatMessage(hwnd syscall.Handle, message string) bool {
	// 激活窗口
	setForegroundWindow.Call(uintptr(hwnd))
	time.Sleep(300 * time.Millisecond)

	// 设置剪贴板内容
	if !SetClipboardText(message) {
		return false
	}

	// 模拟Ctrl+V粘贴
	// Ctrl down
	keybdEvent(VK_CONTROL, 0, 0, 0)
	time.Sleep(50 * time.Millisecond)

	// V down
	keybdEvent(VK_V, 0, 0, 0)
	time.Sleep(50 * time.Millisecond)

	// V up
	keybdEvent(VK_V, 0, KEYEVENTF_KEYUP, 0)
	time.Sleep(50 * time.Millisecond)

	// Ctrl up
	keybdEvent(VK_CONTROL, 0, KEYEVENTF_KEYUP, 0)
	time.Sleep(50 * time.Millisecond)

	// 模拟Enter发送
	SendKey(hwnd, VK_RETURN)
	return true
}

func keybdEvent(bVk byte, bScan byte, dwFlags uint32, dwExtraInfo uintptr) {
	procKeybdEvent.Call(
		uintptr(bVk),
		uintptr(bScan),
		uintptr(dwFlags),
		dwExtraInfo,
	)
}
