// +build linux darwin

package di

import (
	appcore "github.com/fernandochristyanto/retsgo/app/core"
	appdb "github.com/fernandochristyanto/retsgo/app/db"
	v1handler "github.com/fernandochristyanto/retsgo/app/handlers/v1"
	appmiddleware "github.com/fernandochristyanto/retsgo/app/middleware"
	apprepo "github.com/fernandochristyanto/retsgo/repositories"
	appservice "github.com/fernandochristyanto/retsgo/services"
	"github.com/google/wire"
)

/**
 * HELPERS
 */
func GetLogger() *appcore.Log {
	wire.Build(appcore.NewLogrusLogger, appcore.NewLogger)
	return &appcore.Log{}
}

/**
 * MIDDLEWARES
 */
func GetLoggerMiddleware() *appmiddleware.LogMiddleware {
	wire.Build(GetLogger, appmiddleware.NewLogMiddleware)
	return &appmiddleware.LogMiddleware{}
}
func GetPanicMiddleware() *appmiddleware.PanicMiddleware {
	wire.Build(GetLogger, appmiddleware.NewPanicMiddleware)
	return &appmiddleware.PanicMiddleware{}
}

/**
 * REPOSITORIES & SERVICES
 */
func GetTodoRepository() *apprepo.TodoRepository {
	wire.Build(appdb.GetDB, apprepo.NewTodoRepository)
	return &apprepo.TodoRepository{}
}

func GetTodoService() *appservice.TodoService {
	wire.Build(GetLogger, GetTodoRepository, appservice.NewTodoService)
	return &appservice.TodoService{}
}

/**
 * HANDLERS
 */
func GetTodoHandler() *v1handler.TodoHandler {
	wire.Build(GetTodoService, v1handler.NewTodoHandler)
	return &v1handler.TodoHandler{}
}
