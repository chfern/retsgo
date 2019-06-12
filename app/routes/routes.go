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

	/**
	 * V1 APIs
	 */
	versionone := r.PathPrefix("/v1").Subrouter()

	// Auth Routes
	AddRoutes(versionone, func(m *Mapper) {
		jwtHandler := di.GetJWTHandler()
		m.Map("/token").To(jwtHandler.Login).Methods("POST")
	})
	return *r
}
