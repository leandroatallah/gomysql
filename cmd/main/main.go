package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leandroatallah/gomysql/pkg/config"
	"github.com/leandroatallah/gomysql/pkg/controllers"
	"github.com/leandroatallah/gomysql/pkg/models"
	"github.com/leandroatallah/gomysql/pkg/routes"
)

func main() {
	db := config.Connect()
	userModel := &models.UserModel{DB: db}
	userController := &controllers.UserController{Model: userModel}

	r := mux.NewRouter()
	routes.RegisterRoutes(r, userController)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3333", r))
}
