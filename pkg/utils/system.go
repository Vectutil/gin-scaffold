package utils

import "golang.org/x/crypto/bcrypt"

func GetHashStr(str string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
}
