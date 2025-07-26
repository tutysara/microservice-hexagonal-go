package domain

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// this adapter should implement secondary port (CustomerRepository)
type CustomerRepositoryDb struct {
	client *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "banking"
)

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := c.client.Query(findAllSql)

	if err != nil {
		log.Print("Error while querying customer table: " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Print("Error while scanning customers: " + err.Error())
			return nil, err

		}
		customers = append(customers, c)
	}

	return customers, nil
}

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("pgx", connStr)
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
	client, err := connectDB()
	if err != nil {
		panic(err)
	}
	return CustomerRepositoryDb{client}
}
