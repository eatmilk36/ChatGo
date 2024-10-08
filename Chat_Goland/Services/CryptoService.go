package Services

import (
	"crypto/md5"
	"encoding/hex"
)

type CryptoService struct{}

func (c *CryptoService) Md5Hash(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	md5Hash := hash.Sum(nil)

	return hex.EncodeToString(md5Hash)
}
