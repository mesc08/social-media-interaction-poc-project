package endpoints

import (
	"user_post_creation/utils"

	"github.com/gorilla/mux"
)

func CreateEndpoints(router *mux.Router) {
	router.HandleFunc("/user/login", utils.Login).Methods("POST")
	router.HandleFunc("/user/signup", utils.Signup).Methods("POST")
	router.HandleFunc("/user/{id}", utils.UserDetails).Methods("GET")
	router.HandleFunc("/posts", utils.CreatePost).Methods("GET")
	router.HandleFunc("/posts/{id}", utils.GetPostById).Methods("GET")
}
