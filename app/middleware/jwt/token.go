package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/fernandochristyanto/retsgo/config"
	"time"
)

func GenerateToken(user *models.User) string {
	conf := config.GetConf()

	expirationTime := time.Now().Add(time.Duration(conf.JWTTimeoutMins) * time.Minute)
	claims := &Claims{
		Username: user.Username,
		ID:       int64(user.ID),
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(conf.JWTSecretKey))

	if err != nil {
		panic(err)
	}
	return tokenStr
}
