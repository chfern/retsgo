package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims create a struct that will be encoded to a JWT.
type Claims struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
