// Package  provides
package main

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "server.com/crud/controllers"
	middleware "server.com/crud/middlewares"
)

func main() {
	app := mux.NewRouter()

	app.Use(middleware.Logger)

	app.HandleFunc("/", controller.GetContacts).Methods("GET")
	app.HandleFunc("/add-contact", controller.AddContact).Methods("POST")
	app.HandleFunc("/remove", controller.RemoveContact).Methods("POST")

	http.ListenAndServe(":8000", app)
}
