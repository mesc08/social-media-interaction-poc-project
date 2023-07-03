package main

import (
	"net/http"
	"user_post_creation/endpoints"

	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	router := mux.NewRouter()
	endpoints.CreateEndpoints(router)
	http.ListenAndServe(":8100", router)
}
