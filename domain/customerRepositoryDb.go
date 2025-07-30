package domain

import (
	"database/sql"
	"fmt"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "banking"
)

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=$1"
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("Error while querying customer table " + err.Error())
	}

	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers) //LRN: TODO: GO compiler didn't protected from passing customer instead of pointer, was a runtime error
	if err != nil {
		logger.Error("Error while scanning customers: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error in server")

	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=$1"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
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

func connectDB() (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := connectDB() // TODO: how to create connection only when used?
	if err != nil {
		panic(err)
	}
	return CustomerRepositoryDb{client}
}
