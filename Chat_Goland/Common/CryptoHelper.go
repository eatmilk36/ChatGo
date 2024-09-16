package Common

import (
	"crypto/md5"
	"encoding/hex"
)

type CryptoHelper struct{}

func (c *CryptoHelper) Md5Hash(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	md5Hash := hash.Sum(nil)

	return hex.EncodeToString(md5Hash)
}
