package utils

import "errors"

var (
	// ErrUserNotFound 用户未找到错误
	ErrUserNotFound = errors.New("用户未找到")
	// ErrInvalidUserInfo 无效的用户信息错误
	ErrInvalidUserInfo = errors.New("无效的用户信息")
) 