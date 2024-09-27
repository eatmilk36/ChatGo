package Interface

import (
	"Chat_Goland/Services"
)

type JwtServiceInterface interface {
	GenerateJWT(username string) (string, error)

	ValidateJWT(tokenString string) (*Services.MyCustomClaims, error)
}
