package model

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePasswordHash(pwd string) string {
	return Md5(pwd)
}

func Md5(pwd string) string {
	hasher := md5.New()
	hash := hasher.Sum([]byte(pwd))
	return hex.EncodeToString(hash)
}
