package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leandroatallah/gomysql/pkg/models"
)

func HTTPInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("[ERROR] %s %s: %v\n", r.Method, r.URL.Path, err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userPayload models.User
	json.NewDecoder(r.Body).Decode(&userPayload)
	user, err := models.CreateUser(userPayload)
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	if user.Id == 0 {
		w.WriteHeader(http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.UserToUserResponse(user))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	user, err := models.GetUserById(userID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserToUserResponse(user))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	user, err := models.DeleteUser(userID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserToUserResponse(user))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	found, err := models.GetUserById(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var userPayload models.User
	json.NewDecoder(r.Body).Decode(&userPayload)
	found.Username = userPayload.Username
	found.Password = userPayload.Password
	user, err := models.UpdateUser(found)
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserToUserResponse(user))
}
