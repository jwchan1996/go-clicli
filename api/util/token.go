package util

import (
	"encoding/base64"
	"log"
)

func CreateToken(uname string) string {
	newId := Cipher(uname+"acgzone2333@132SAMA")
	token := base64.StdEncoding.EncodeToString([]byte(newId))
	return token
}

func ResolveToken(token string) string {
	str, _ := base64.StdEncoding.DecodeString(token)
	log.Printf("%s", str)
	res := string(str)
	return res
}
