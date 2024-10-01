package Interface

import (
	"Chat_Goland/Services"
)

type JwtServiceInterface interface {
	GenerateJWT(username string, id int) (string, error)

	ValidateJWT(tokenString string) (*Services.MyCustomClaims, error)
}
