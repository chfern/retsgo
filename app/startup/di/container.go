// +build linux darwin

package di

import (
	"github.com/fernandochristyanto/retsgo/app/core"
	appdb "github.com/fernandochristyanto/retsgo/app/db"
	v1handler "github.com/fernandochristyanto/retsgo/app/handlers/v1"
	appmiddleware "github.com/fernandochristyanto/retsgo/app/middleware"
	jwtvalidator "github.com/fernandochristyanto/retsgo/app/middleware/jwt/validator"
	apprepo "github.com/fernandochristyanto/retsgo/repositories"
	appservice "github.com/fernandochristyanto/retsgo/services"
	"github.com/google/wire"
)

/**
 * HELPERS
 */
func GetLogger() *core.Log {
	panic(wire.Build(core.NewLogrusLogger, core.NewLogger))
}

/**
 * REPOSITORIES
 */
func GetUserRepository() *apprepo.UserRepository {
	panic(wire.Build(appdb.GetDB, apprepo.NewUserRepository))
}

/**
 * SERVICES
 */
func GetUserService() *appservice.UserService {
	panic(wire.Build(GetUserRepository, appservice.NewUserService))
}

/**
 * MIDDLEWARES
 */
func GetLoggerMiddleware() *appmiddleware.LogMiddleware {
	panic(wire.Build(GetLogger, appmiddleware.NewLogMiddleware))
}

func GetPanicMiddleware() *appmiddleware.PanicMiddleware {
	panic(wire.Build(GetLogger, appmiddleware.NewPanicMiddleware))
}

func GetJWTMiddleware() *appmiddleware.JWTMiddleware {
	panic(wire.Build(GetUserService, appmiddleware.NewJWTMiddleware))
}
func GetJWTUserValidator() *jwtvalidator.UserValidator {
	panic(wire.Build(GetUserService, jwtvalidator.NewUserValidator))
}

/**
 * HANDLERS
 */
func GetJWTHandler() *v1handler.JWTHandler {
	panic(wire.Build(GetUserService, v1handler.NewJWTHandler))
}
