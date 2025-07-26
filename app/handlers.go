package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/tutysara/banking-go/app/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

// this adapter (REST handlers) depends on the primary port aka service interface
type CustomerHandlers struct {
	service service.DefaultCustomerService
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post request received")
}

func (h CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := h.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
