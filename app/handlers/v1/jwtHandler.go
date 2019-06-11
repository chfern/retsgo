package v1

import (
	appjwt "github.com/fernandochristyanto/retsgo/app/middleware/jwt"
	apprepo "github.com/fernandochristyanto/retsgo/repositories"
	"github.com/json-iterator/go"
	"net/http"
)

type JWTHandler struct {
	userRepository *apprepo.UserRepository
}

func (handler JWTHandler) Login(w http.ResponseWriter, r *http.Request) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	var credentials appjwt.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		panic(err)
	}

	user := handler.userRepository.GetByUsernameAndPassword(
		credentials.Username,
		credentials.Password,
	)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	// Create the JWT token
	tokenStr := appjwt.GenerateToken(user)
	if err != nil {
		panic(err)
	}
	resp, _ := json.Marshal(map[string]string{
		"token": tokenStr,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func NewJWTHandler(userRepository *apprepo.UserRepository) *JWTHandler {
	return &JWTHandler{
		userRepository: userRepository,
	}
}
