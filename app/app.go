package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tutysara/banking-go/domain"
	"github.com/tutysara/banking-go/service"
)

func sanityCheck() {
	if os.Getenv("DBHOST") == "" {
		panic("DBHOST is undefined")
	}
	if os.Getenv("DBPORT") == "" {
		panic("DBPORT is undefined")
	}
	if os.Getenv("DBUSER") == "" {
		panic("DBUSER is undefined")
	}
	if os.Getenv("DBPASSWORD") == "" { // LRN: TODO: Fatal log also exits the application, this is weird
		log.Println("Warn: DBPASSWORD is undefined")
	}
	if os.Getenv("DBNAME") == "" {
		panic("DBNAME is undefined")
	}
	if os.Getenv("SERVER_ADDRESS") == "" {
		panic("SERVER_ADDRESS is undefined")
	}
	if os.Getenv("SERVER_PORT") == "" {
		panic("SERVER_PORT is undefined")
	}

}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	// wiring the application
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	// starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), router))
}

//SERVER_ADDRESS=localhost SERVER_PORT=8080 DBHOST=localhost DBPORT=5432 DBUSER=postgres DBNAME=banking go run main.go
