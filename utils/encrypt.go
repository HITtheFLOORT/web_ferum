package utils

import (
	"crypto/md5"
	"encoding/hex"
)
const secret ="tangminghao"
func EncryptPassword(o string)string{
	h:=md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(o)))
}