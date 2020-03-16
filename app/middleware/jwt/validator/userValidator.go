package validator

import (
	appjwt "github.com/fernandochristyanto/retsgo/app/middleware/jwt"
	appservice "github.com/fernandochristyanto/retsgo/services"
)

type UserValidator struct {
	userService *appservice.UserService
}

func (validator UserValidator) Validate(claims *appjwt.Claims) bool {
	if claims.Role != "user" {
		return false
	}

	user := validator.userService.GetByID(claims.ID)
	if user == nil {
		return false
	}

	return true
}

func NewUserValidator(userService *appservice.UserService) *UserValidator {
	return &UserValidator{
		userService: userService,
	}
}
