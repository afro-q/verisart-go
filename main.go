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
	middleware "github.com/quinlanmorake/verisart-go/middleware"
	
	userHttpHandlers "github.com/quinlanmorake/verisart-go/user/handlers/http"
	certifcateHttpHandlers "github.com/quinlanmorake/verisart-go/certificate/handlers/http"
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

	// Authentication onwards
	userSubRoutes := router.PathPrefix("/users/{userId}/certificates").Subrouter()
	userSubRoutes.Use(middleware.Authentication)
	userSubRoutes.HandleFunc("/", certifcateHttpHandlers.LoadCertificatesForUser).Methods(coreTypes.HTTP_GET) // Load all certificates for user

	certificateRoutes := router.PathPrefix("/certificates").Subrouter()
	certificateRoutes.Use(middleware.Authentication)

	certificateRoutes.HandleFunc("/", certifcateHttpHandlers.Create).Methods(coreTypes.HTTP_POST) // Create a certificate
	certificateRoutes.HandleFunc("/{certificateId}", certifcateHttpHandlers.Update).Methods(coreTypes.HTTP_PUT) // Update a certificate
	certificateRoutes.HandleFunc("/{certificateId}", certifcateHttpHandlers.Delete).Methods(coreTypes.HTTP_DELETE) // Delete a certificate

	certificateRoutes.HandleFunc("/{certificateId}/transfers", certifcateHttpHandlers.CreateTransfer).Methods(coreTypes.HTTP_POST) // Create a transfer
	certificateRoutes.HandleFunc("/{certificateId}/transfers", certifcateHttpHandlers.AcceptTransfer).Methods(coreTypes.HTTP_PUT)  // Accept a transfer

	http.Handle("/", router)

	log.Printf("Application listening on %v\n", config.AppConfig.Server.GetListenAddress())
	if err := http.ListenAndServe(config.AppConfig.Server.GetListenAddress(), nil); err != nil {
		log.Fatalf("%v", err)
	}
}

func handleInitializeError(error coreTypes.Result) {
	log.Fatalf("%#v \n", error) // I want to print out the whole message
}
