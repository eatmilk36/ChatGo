package Ineterface

import "chat/Common"

type JwtInterface interface {
	GenerateJWT(username string) (string, error)

	ValidateJWT(tokenString string) (*Common.MyCustomClaims, error)
}
