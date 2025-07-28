package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutysara/banking-go/service"
)

// this adapter (REST handlers) depends on the primary port aka service interface
type CustomerHandlers struct {
	service service.DefaultCustomerService
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post request received")
}

func (h CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := h.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, customers)

	}
}

func (h CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	log.Println("id", id)
	customer, err := h.service.GetCustomer(id)
	log.Println("customer=", customer)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil { // Encode works even with pointer to struct
		panic(err)
	}
}
