package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user_post_creation/model"

	"github.com/gorilla/mux"
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

func UserDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid := params["id"]
	userDetails, err := GetUserDetails(userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDetails)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	post, err := PostById(idStr)
	if err != nil {
		http.Error(w, "Invalid post id", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := model.Post{}
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultId, err := SavePost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Id = resultId
	response := model.Response{Data: post, Status: 200}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
