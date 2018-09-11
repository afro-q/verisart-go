package main

import (
	"log"
	"net/http"

	gorillaMux "github.com/gorilla/mux"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	authentication "github.com/quinlanmorake/verisart-go/authentication"
	config "github.com/quinlanmorake/verisart-go/config"
	database "github.com/quinlanmorake/verisart-go/database"
	users "github.com/quinlanmorake/verisart-go/user"

	userHttpHandlers "github.com/quinlanmorake/verisart-go/user/handlers/http"

	createHandlers "github.com/quinlanmorake/verisart-go/create/handlers"
	deleteHandlers "github.com/quinlanmorake/verisart-go/delete/handlers"
	listHandlers "github.com/quinlanmorake/verisart-go/list/handlers"
	middleware "github.com/quinlanmorake/verisart-go/middleware"
	transferHandlers "github.com/quinlanmorake/verisart-go/transfer/handlers"
	updateHandlers "github.com/quinlanmorake/verisart-go/update/handlers"
)

func main() {
	if loadConfigResult := config.LoadConfig(); loadConfigResult.IsNotOk() {
		handleInitializeError(loadConfigResult)
	}

	if initializeDbResult := database.Init(config.AppConfig); initializeDbResult.IsNotOk() {
		handleInitializeError(initializeDbResult)
	}

	if initializeAuthResult := authentication.Init(config.AppConfig); initializeAuthResult.IsNotOk() {
		handleInitializeError(initializeAuthResult)
	}

	// Before anything else, lets add some users to the database
	if addUsersResult := users.Init(); addUsersResult.IsNotOk() {
		handleInitializeError(addUsersResult)
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

	// Let us load all users and generate a certificate without being authenticated
	router.HandleFunc("/users", userHttpHandlers.LoadAllUsers).Methods(coreTypes.HTTP_GET)
	router.HandleFunc("/users/{userId}/token", userHttpHandlers.GenerateToken).Methods(coreTypes.HTTP_GET)

	userSubRoutes := router.PathPrefix("/users/{userId}/certificats").Subrouter()
	userSubRoutes.Use(middleware.Authentication)
	userSubRoutes.HandleFunc("/", listHandlers.HandleHttpRequest).Methods(coreTypes.HTTP_GET)

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

func handleInitializeError(error coreTypes.Result) {
	log.Fatalf("%#v \n", error) // I want to print out the whole message
}
