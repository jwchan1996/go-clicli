package util

import (
	"encoding/base64"
)

func CreateToken(uname string,  role string) string {
	newId := Cipher(uname + role + "clicli2333@132SAMA#ojbk")
	token := base64.StdEncoding.EncodeToString([]byte(newId))
	return token
}
