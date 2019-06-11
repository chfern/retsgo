package routes

import (
	"github.com/fernandochristyanto/retsgo/app/build/di"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) mux.Router {
	r = mux.NewRouter().StrictSlash(true)

	/**
	 * Declare middleware and middleware dependencies here
	 */
	jwtMiddleware := di.GetJWTMiddleware()
	userValidator := di.GetJWTUserValidator().Validate

	/**
	 * V1 APIs
	 */
	versionone := r.PathPrefix("/v1").Subrouter()

	// Auth Routes
	AddRoutes(versionone, func(m *Mapper) {
		jwtHandler := di.GetJWTHandler()
		m.Map("/token").To(jwtHandler.Login).Methods("POST")
	})

	// Todo Routes
	AddRoutes(versionone.PathPrefix("/todos").Subrouter(), func(m *Mapper) {
		todoHandler := di.GetTodoHandler()
		m.Map("").To(todoHandler.GetAll).Methods("GET")
		m.Map("/{id}").To(todoHandler.GetByID).Methods("GET")
	})

	// User Routes
	AddRoutes(versionone.PathPrefix("/users").Subrouter(), func(m *Mapper) {
		userHandler := di.GetUserHandler()
		m.Map("/{id}/todos").WithMiddleware(jwtMiddleware.With(userValidator).M).To(userHandler.GetTodos).Methods("GET")
	})
	return *r
}
