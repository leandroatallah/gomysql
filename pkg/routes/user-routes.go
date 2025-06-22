package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/leandroatallah/gomysql/pkg/controllers"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

const maxBodySize = 1 << 20 // 1MB

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			f(w, r)
		}
	}
}

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

func LimitBodySize() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f
}

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/users", Chain(controllers.GetAllUsers, Method("GET"), LimitBodySize(), Logging())).Methods("GET")
	r.HandleFunc("/users", Chain(controllers.CreateUser, Method("POST"), LimitBodySize(), Logging())).Methods("POST")
	r.HandleFunc("/users/{id}", Chain(controllers.GetUserById, Method("GET"), LimitBodySize(), Logging())).Methods("GET")
	r.HandleFunc("/users/{id}", Chain(controllers.DeleteUser, Method("DELETE"), LimitBodySize(), Logging())).Methods("DELETE")
	r.HandleFunc("/users/{id}", Chain(controllers.UpdateUser, Method("PUT"), LimitBodySize(), Logging())).Methods("PUT")
}
