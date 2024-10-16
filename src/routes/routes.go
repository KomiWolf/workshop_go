package routes

import (
	"net/http"
	"poc-workshop-go/controllers"
	"poc-workshop-go/middlewares"

	"github.com/gorilla/mux"
)

// Init sets up all the routes for the server
// it could be splited up in differents sub fonctions
func Init(r *mux.Router) {

	r.HandleFunc("/", controllers.Welcome).Methods("GET")
	r.HandleFunc("/hello", controllers.Hello).Methods(("GET"))
	r.Use(middlewares.LoggerMiddleware)

	helloHandler := http.HandlerFunc(controllers.Hello)
	r.HandleFunc("/auth/hello", middlewares.Auth(helloHandler)).Methods("GET")
	r.HandleFunc("/whoami/{user}", controllers.Whoami).Methods("GET")
	r.HandleFunc("/add", controllers.AddUser).Methods("POST")
	r.HandleFunc("/get/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/del/{id}", controllers.DeleteUser).Methods("DELETE")
}
