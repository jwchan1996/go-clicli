package util

import (
	"encoding/base64"
)

func CreateToken(uqq int) string {
	newId := Cipher(string(uqq))
	token := base64.StdEncoding.EncodeToString([]byte(newId))
	return token
}

func ResolveToken(token string) string {
	str, _ := base64.StdEncoding.DecodeString(token)
	res := string(str)
	return res
}
