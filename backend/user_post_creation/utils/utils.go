package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user_post_creation/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	valid, err := ValidateUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, fmt.Errorf("User not found or password mismatched").Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "User logged in successfully")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := RegisterUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "User logged in successfully")

}
