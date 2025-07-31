package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutysara/banking-go/dto"
	"github.com/tutysara/banking-go/service"
)

// this is the primary adapter that uses primary port
type AccountHandlers struct {
	service service.AccountService
}

func NewAccountHandlers(s service.AccountService) AccountHandlers {
	return AccountHandlers{
		service: s,
	}
}

func (h AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
