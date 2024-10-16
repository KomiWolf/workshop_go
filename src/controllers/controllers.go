package controllers

import (
	"encoding/json"
	"net/http"
	model "poc-workshop-go/models"
	"strconv"

	"github.com/gorilla/mux"
)

// Welcome serves an endpoint and returns a base welcome message to the client
func Welcome(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode("welcome to the workshop !"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode("Hello, World!"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Whoami(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := json.NewEncoder(w).Encode("I am " + vars["user"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var usr model.Users

	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := model.Creatuser(usr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode("msg: The user is successfully created."); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := model.Getuser(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = model.Deluser(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode("msg: The user is deleted successfully"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
