package main

import (
	"log"
	"net/http"

	gorillaMux "github.com/gorilla/mux"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	authentication "github.com/quinlanmorake/verisart-go/authentication"	
	config "github.com/quinlanmorake/verisart-go/config"
	database "github.com/quinlanmorake/verisart-go/database"

	createHandlers "github.com/quinlanmorake/verisart-go/create/handlers"
	deleteHandlers "github.com/quinlanmorake/verisart-go/delete/handlers"
	listHandlers "github.com/quinlanmorake/verisart-go/list/handlers"
	middleware "github.com/quinlanmorake/verisart-go/middleware"
	transferHandlers "github.com/quinlanmorake/verisart-go/transfer/handlers"
	updateHandlers "github.com/quinlanmorake/verisart-go/update/handlers"
)

func main() {
	if loadConfigResult := config.LoadConfig(); loadConfigResult.IsNotOk() {
		log.Fatalf("%#v \n", loadConfigResult) // I want to print out the whole message
	}

	if initializeDbResult := database.Init(config.AppConfig); initializeDbResult.IsNotOk() {
		log.Fatalf("%#v \n", initializeDbResult)
	}

	if initializeAuthResult := authentication.Init(config.AppConfig); initializeAuthResult.IsNotOk() {
		log.Fatalf("%#v \n", initializeAuthResult)
	}
	
	/*
	   It is not an overly complicated thing to create one's own router and minimize
	   the dependency on other libraries; in this case there would have been a little
	   work to do around the pulling out of parameters, which I opted not to do so as
	   to save time.
	*/
	router := gorillaMux.NewRouter()

	// Add middleware that is applied across board
	router.Use(middleware.CORS)
	router.Use(middleware.ContentType)

	userRoutes := router.PathPrefix("/users").Subrouter()
	userRoutes.Use(middleware.Authentication)
	userRoutes.HandleFunc("/{userId}/certificates", listHandlers.HandleHttpRequest).Methods(coreTypes.HTTP_GET)
	
	certificateRoutes := router.PathPrefix("/certificates").Subrouter()
	certificateRoutes.Use(middleware.Authentication)

	certificateRoutes.HandleFunc("/", createHandlers.HandleHttpRequest).Methods(coreTypes.HTTP_POST)
	certificateRoutes.HandleFunc("/{certificateId}", updateHandlers.HandleHttpRequest).Methods(coreTypes.HTTP_PUT)
	certificateRoutes.HandleFunc("/{certificateId}", deleteHandlers.HandleHttpRequest).Methods(coreTypes.HTTP_DELETE)

	certificateRoutes.HandleFunc("/{certificateId}/transfers", transferHandlers.HandleCreateTransferHttpRequest).Methods(coreTypes.HTTP_POST)
	certificateRoutes.HandleFunc("/{certificateId}/transfers", transferHandlers.HandleAcceptTransferHttpRequest).Methods(coreTypes.HTTP_PUT)

	http.Handle("/", router)

	log.Printf("Application listening on %v\n", config.AppConfig.Server.GetListenAddress())
	if err := http.ListenAndServe(config.AppConfig.Server.GetListenAddress(), nil); err != nil {
		log.Fatalf("%v", err)
	}
}
