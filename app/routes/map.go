package routes

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

type Mapper struct {
	path        string
	router      *mux.Router
	middlewares []alice.Constructor
}

func (m *Mapper) Map(path string) *Mapper {
	return &Mapper{
		path:   path,
		router: m.router,
	}
}

func (m *Mapper) WithMiddleware(constructors ...alice.Constructor) *Mapper {
	m.middlewares = constructors
	return m
}

func (m *Mapper) To(handler func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	defaultMiddlewares := GetDefaultMiddlewares()
	return m.router.Handle(m.path, alice.New(defaultMiddlewares...).Append(m.middlewares...).Then(http.HandlerFunc(handler)))
}
