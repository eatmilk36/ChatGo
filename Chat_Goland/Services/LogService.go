package Services

import (
	"crypto/md5"
	"encoding/hex"
)

type LogService struct{}

func (c *LogService) LogError(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	md5Hash := hash.Sum(nil)

	return hex.EncodeToString(md5Hash)
}
