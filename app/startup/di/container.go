// +build linux darwin

package di

import (
	appcore "github.com/fernandochristyanto/retsgo/app/core"
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
func GetLogger() *appcore.Log {
	panic(wire.Build(appcore.NewLogrusLogger, appcore.NewLogger))
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
	panic(wire.Build(GetUserRepository, appmiddleware.NewJWTMiddleware))
}

func GetJWTUserValidator() *jwtvalidator.UserValidator {
	panic(wire.Build(GetUserRepository, jwtvalidator.NewUserValidator))
}

/**
 * REPOSITORIES & SERVICES
 */
func GetTodoService() *appservice.TodoService {
	panic(wire.Build(appdb.GetDB, appservice.NewTodoService))
}

func GetTodoRepository() *apprepo.TodoRepository {
	panic(wire.Build(GetTodoService, apprepo.NewTodoRepository))
}

func GetUserService() *appservice.UserService {
	panic(wire.Build(appdb.GetDB, appservice.NewUserService))
}

func GetUserRepository() *apprepo.UserRepository {
	panic(wire.Build(GetUserService, apprepo.NewUserRepository))
}

/**
 * HANDLERS
 */
func GetTodoHandler() *v1handler.TodoHandler {
	panic(wire.Build(GetTodoRepository, v1handler.NewTodoHandler))
}

func GetJWTHandler() *v1handler.JWTHandler {
	panic(wire.Build(GetUserRepository, v1handler.NewJWTHandler))
}

func GetUserHandler() *v1handler.UserHandler {
	panic(wire.Build(GetUserRepository, v1handler.NewUserHandler))
}
