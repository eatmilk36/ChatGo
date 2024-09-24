package Mock

import (
	"Chat_Goland/Services"
	"github.com/stretchr/testify/mock"
)

type Jwt struct {
	mock.Mock
}

func (j *Jwt) GenerateJWT(username string) (string, error) {
	args := j.Called(username)
	return args.String(0), args.Error(1)
}

func (j *Jwt) ValidateJWT(tokenString string) (*Services.MyCustomClaims, error) {
	args := j.Called(tokenString)
	if claims := args.Get(0); claims != nil {
		return claims.(*Services.MyCustomClaims), args.Error(1)
	}
	return nil, args.Error(1)
}
