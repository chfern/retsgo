package routes

import (
	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router, mfunc func(m *Mapper)) {
	mapper := Mapper{
		router: router,
	}
	mfunc(&mapper)
}
