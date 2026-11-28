package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	md5_ := md5.New()
	md5_.Write([]byte(str))
	cipherStr := md5_.Sum(nil)
	md5Str := hex.EncodeToString(cipherStr)

	return md5Str
}
