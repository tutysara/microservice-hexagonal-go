package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutysara/banking-go/domain"
	"github.com/tutysara/banking-go/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring the application
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
