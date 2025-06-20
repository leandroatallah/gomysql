package routes

import (
	"github.com/gorilla/mux"
	"github.com/leandroatallah/gomysql/pkg/controllers"
)

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
}
