package Mock

import "github.com/stretchr/testify/mock"

type CryptoHelper struct {
	mock.Mock
}

func (c *CryptoHelper) Md5Hash(value string) string {
	args := c.Called(value)
	return args.String(0)
}
