package middleware

import (
	"fmt"
	"github.com/fernandochristyanto/retsgo/app/core"
	templateapi "github.com/fernandochristyanto/retsgo/app/core/templates/api"
	"github.com/json-iterator/go"
	"net/http"
)

type PanicMiddleware struct {
	logger *core.Log
}

func (panicMiddleware PanicMiddleware) M(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json := jsoniter.ConfigCompatibleWithStandardLibrary

		defer func() {
			if r := recover(); r != nil {
				w.Header().Set("Content-Type", "application/json")
				httpException, valid := r.(*core.HttpException)

				if valid {
					errJSON, _ := json.Marshal(templateapi.NewError(httpException.Error))
					w.WriteHeader(httpException.Code)
					w.Write(errJSON)
				} else {
					errJSON, _ := json.Marshal(templateapi.NewError(fmt.Sprint(r)))
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(errJSON)
				}
			}
		}()

		handler.ServeHTTP(w, r) // Possibly error thrown
	})
}

func NewPanicMiddleware(logger *core.Log) *PanicMiddleware {
	return &PanicMiddleware{
		logger: logger,
	}
}
