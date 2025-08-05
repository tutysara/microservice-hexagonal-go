package domain

import (
	"database/sql"
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

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	findSql := `select customer_id, opening_date, account_type, amount, status 
	from accounts
	where account_id=$1`

	var account Account
	err := d.client.Get(&account, findSql, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Account with the Id found: " + err.Error())
			return nil, errs.NewNotFoundError("No account found with id :" + accountId)
		} else {
			logger.Error("Error while scanning account: " + err.Error())
			return nil, errs.NewUnexpectedError("Error while scanning account: " + err.Error())
		}
	}
	return &account, nil

}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// insert into transaction
	insertTransactionSql := `insert into transactions (account_id, amount, transaction_type, transaction_date)
	values ($1, $2, $3, $4)
	RETURNING transaction_id
	`
	result := tx.QueryRow(insertTransactionSql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// get the last inserted transactionId
	var transactionId int64
	err = result.Scan(&transactionId)

	if err != nil {
		logger.Error("Error while getting transactionId: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	if t.IsWithDrawal() {
		_, err = tx.Exec("UPDATE Accounts SET amount= amount-$1 where account_id=$2", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE Accounts SET amount= amount+$1 where account_id=$2", t.Amount, t.AccountId)

	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// if everything is ok commit the transaction
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// get lastest accounts info from table
	account, appErr := d.FindBy(t.AccountId)
	if err != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{
		client: client,
	}
}
