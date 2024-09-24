package Ineterface

import (
	"Chat_Goland/Services"
)

type JwtInterface interface {
	GenerateJWT(username string) (string, error)

	ValidateJWT(tokenString string) (*Services.MyCustomClaims, error)
}
