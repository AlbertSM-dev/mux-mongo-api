package main

import (
	"log"
	"mux-mongo-api/configs"
	"mux-mongo-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":7000", router))
}
