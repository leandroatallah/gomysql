package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/leandroatallah/gomysql/pkg/models"
)

func HTTPInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("[ERROR] %s %s: %v\n", r.Method, r.URL.Path, err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

type UserResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func UserToUserResponse(user models.User) UserResponse {
	var userRes UserResponse
	userRes.Username = user.Username
	userRes.CreatedAt = user.CreatedAt
	return userRes
}

type UserController struct {
	Model *models.UserModel
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Model.GetAllUsers()
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userPayload models.User
	json.NewDecoder(r.Body).Decode(&userPayload)
	user, err := c.Model.CreateUser(userPayload)
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	if user.Id == 0 {
		w.WriteHeader(http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserToUserResponse(user))
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	user, err := c.Model.GetUserById(userID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserToUserResponse(user))
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	user, err := c.Model.DeleteUser(userID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserToUserResponse(user))
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	found, err := c.Model.GetUserById(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var userPayload models.User
	json.NewDecoder(r.Body).Decode(&userPayload)
	found.Username = userPayload.Username
	found.Password = userPayload.Password
	user, err := c.Model.UpdateUser(found)
	if err != nil {
		HTTPInternalServerError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserToUserResponse(user))
}
