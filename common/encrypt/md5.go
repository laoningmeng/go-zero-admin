package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(originPass string) string {
	hash := md5.New()
	data := []byte(originPass)
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
