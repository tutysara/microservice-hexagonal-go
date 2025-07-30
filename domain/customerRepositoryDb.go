package domain

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/tutysara/banking-go/errs"
	"github.com/tutysara/banking-go/logger"
)

// this adapter should implement secondary port (CustomerRepository)
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = ""
// 	dbname   = "banking"
// )

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error

	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=$1"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("Error while querying customer table " + err.Error())
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=$1"
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error no customers returned " + err.Error())
			return nil, errs.NewNotFoundError("No Customer Found in Repo")
		} else {
			logger.Error("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected error in server")
		}

	}
	log.Println("Got from db :", c)
	return &c, nil
}

func NewCustomerRepositoryDb(dbclient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbclient}
}
