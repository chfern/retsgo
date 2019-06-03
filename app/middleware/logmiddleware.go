package middleware

import (
	"github.com/fernandochristyanto/retsgo/app/core"
	"net/http"
)

type LogMiddleware struct {
	logger *core.Log
}

func (logMiddleware LogMiddleware) M(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logMiddleware.logger.Infof("%s\n",
			r.Method,
			r.URL,
			r.Proto,
		)
		logMiddleware.logger.Infof("Host: %s\n",
			r.Host,
		)
		for name, headers := range r.Header {
			for _, header := range headers {
				logMiddleware.logger.Infof("%s\n", name, header)
			}
		}
		handler.ServeHTTP(w, r)
	})
}

func NewLogMiddleware(logger *core.Log) *LogMiddleware {
	return &LogMiddleware{
		logger: logger,
	}
}
