package routes

import (
	"github.com/gorilla/mux"

	"API_GO/src/controllers"
)

func SetupRouter(router *mux.Router) {
	//Rotas User
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{name}", controllers.GetUserName).Methods("GET")
}
