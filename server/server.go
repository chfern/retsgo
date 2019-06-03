package main

import (
	"fmt"
	"github.com/fernandochristyanto/retsgo/app/db"
	"github.com/fernandochristyanto/retsgo/app/routes"
	"github.com/fernandochristyanto/retsgo/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	conf := config.GetConf()
	db.InitDB()

	var r mux.Router
	r = routes.RegisterRoutes(&r)

	srv := &http.Server{
		Handler:      &r,
		Addr:         fmt.Sprintf(":%s", conf.General.AppPort),
		WriteTimeout: time.Duration(conf.ListenerWriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(conf.ListenerReadTimeout) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
