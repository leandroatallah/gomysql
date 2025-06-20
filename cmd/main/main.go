package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leandroatallah/gomysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3333", r))
}
