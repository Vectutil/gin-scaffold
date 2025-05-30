package qny

import (
	"crypto/md5"
	"fmt"
	"gin-scaffold/pkg/logger"
)

func EncryptBytes(data []byte) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write(data); err != nil {
		logger.ErrorLogger.Error("encrypt data error")
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func EncryptString(data string) (encrypt string, err error) {
	return EncryptBytes([]byte(data))
}

func MustEncryptString(data string) string {
	result, err := EncryptString(data)
	if err != nil {
		logger.ErrorLogger.Error("encrypt data error")
		return ""
	}
	return result
}
