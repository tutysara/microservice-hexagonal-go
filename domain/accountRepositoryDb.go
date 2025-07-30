package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/tutysara/banking-go/errs"
	"github.com/tutysara/banking-go/logger"
)

// secondary adapter that implements secondary port
type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	insertSql := `insert into accounts (customer_id, opening_date, account_type, amount, status) 
	values ($1, $2, $3, $4, $5)
	RETURNING account_id
	`
	rows := d.client.QueryRowx(insertSql, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	var accountId int64
	err := rows.Scan(&accountId)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(accountId, 10)
	return &a, nil // TODO: should we return a new object with the fields copied? are we doing this for efficiency reason?
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{
		client: client,
	}
}
