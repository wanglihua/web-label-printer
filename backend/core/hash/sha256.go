package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(str string) string {
	sha256_ := sha256.New()
	sha256_.Write([]byte(str))
	cipherStr := sha256_.Sum(nil)
	sha256Str := hex.EncodeToString(cipherStr)

	return sha256Str
}
