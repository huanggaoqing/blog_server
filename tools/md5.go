package tools

import (
	"crypto"
	"encoding/hex"
)

func Md5(s string) string {
	m := crypto.MD5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
