package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

//GetMD5 return an MD5 hashed encoded string
func GetMD5(s string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
