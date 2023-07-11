package main

import (
	"log"
	"net/http"
	"user_post_creation/config"
	"user_post_creation/endpoints"

	"github.com/gorilla/mux"
)

func init() {
	if err := config.ReadConfigFile("config.json"); err != nil {
		log.Fatal("Dropping Service")
	}
}

func main() {
	router := mux.NewRouter()
	endpoints.CreateEndpoints(router)
	http.ListenAndServe(config.ViperConfig.ServiceHost, router)
}
