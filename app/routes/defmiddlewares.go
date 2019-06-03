package routes

import (
	"github.com/fernandochristyanto/retsgo/app/build/di"
	"github.com/justinas/alice"
)

func GetDefaultMiddlewares() []alice.Constructor {
	return []alice.Constructor{
		di.GetLoggerMiddleware().M,
		di.GetPanicMiddleware().M,
	}
}
