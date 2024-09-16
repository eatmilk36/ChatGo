package Ineterface

import "Chat_Goland/Common"

type JwtInterface interface {
	GenerateJWT(username string) (string, error)

	ValidateJWT(tokenString string) (*Common.MyCustomClaims, error)
}
