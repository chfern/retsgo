package middleware

import (
	"github.com/dgrijalva/jwt-go"
	appjwt "github.com/fernandochristyanto/retsgo/app/middleware/jwt"
	"github.com/fernandochristyanto/retsgo/config"
	apprepo "github.com/fernandochristyanto/retsgo/repositories"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	validatorFuncs []func(claims *appjwt.Claims) bool
	userRepository *apprepo.UserRepository
}

func (jwtMiddleware JWTMiddleware) M(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// json := jsoniter.ConfigCompatibleWithStandardLibrary

		// Authenticate jwt token
		authorizationHeader := r.Header.Get("Authorization")
		if strings.Trim(authorizationHeader, "\n") == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// Parse tokenStr
		tokenStrSplit := strings.Split(authorizationHeader, " ")
		if len(tokenStrSplit) != 2 && tokenStrSplit[0] != "Bearer" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		tokenStr := tokenStrSplit[1]
		claims := &appjwt.Claims{}
		secretKey := []byte(config.GetConf().JWTSecretKey)
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(tkn *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			panic(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		isvalid := false
		for _, validatorFunc := range jwtMiddleware.validatorFuncs {
			if validatorFunc(claims) == true {
				isvalid = true
				break
			}
		}
		if !isvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Set Set-Authorization header
		user := jwtMiddleware.userRepository.GetByID(claims.ID)
		w.Header().Set("Set-Authorization", appjwt.GenerateToken(user))

		handler.ServeHTTP(w, r)
	})
}

func (jwtMiddleware JWTMiddleware) With(validator ...func(claims *appjwt.Claims) bool) JWTMiddleware {
	jwtMiddleware.validatorFuncs = validator
	return jwtMiddleware
}

func NewJWTMiddleware(userRepo *apprepo.UserRepository) *JWTMiddleware {
	return &JWTMiddleware{
		userRepository: userRepo,
	}
}
