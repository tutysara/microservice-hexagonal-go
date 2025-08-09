package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// AAA
	// Arrange
	request := TransactionRequest{
		TransactionType: "invalid type",
	}
	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Transaction type should be withdrawal or deposit" {
		t.Error("Invalid message while testing transaction type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// AAA
	// Arrange
	request := TransactionRequest{
		Amount: -23213213,
	}
	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Amount should be > 0" {
		t.Error("Invalid message while testing amount")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing amount")
	}
}
