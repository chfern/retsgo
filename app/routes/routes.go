package routes

import (
	"github.com/fernandochristyanto/retsgo/app/build/di"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) mux.Router {
	r = mux.NewRouter().StrictSlash(true)

	/**
	 * V1 APIs
	 */
	versionone := r.PathPrefix("/v1").Subrouter()
	// todo
	todo := versionone.PathPrefix("/todos").Subrouter()
	AddRoutes(todo, func(m *Mapper) {
		todoHandler := di.GetTodoHandler()
		m.Map("").To(todoHandler.GetAll).Methods("GET")
		m.Map("/{id}").To(todoHandler.GetByID).Methods("GET")
	})
	return *r
}
