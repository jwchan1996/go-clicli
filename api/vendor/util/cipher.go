package util

import (
	"crypto/md5"
	"fmt"
)

func Cipher(pwd string) string {
	str := []byte(pwd)
	res := md5.Sum(str)
	newPwd := fmt.Sprintf("%x", res)
	return newPwd
}
