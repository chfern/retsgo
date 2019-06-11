package validator

import (
	appjwt "github.com/fernandochristyanto/retsgo/app/middleware/jwt"
	apprepo "github.com/fernandochristyanto/retsgo/repositories"
)

type UserValidator struct {
	userRepository *apprepo.UserRepository
}

func (validator UserValidator) Validate(claims *appjwt.Claims) bool {
	if claims.Role != "user" {
		return false
	}

	user := validator.userRepository.GetByID(claims.ID)
	if user == nil {
		return false
	}

	return true
}

func NewUserValidator(userRepo *apprepo.UserRepository) *UserValidator {
	return &UserValidator{
		userRepository: userRepo,
	}
}
