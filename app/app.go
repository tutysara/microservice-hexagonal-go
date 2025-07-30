package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/tutysara/banking-go/domain"
	"github.com/tutysara/banking-go/service"
)

func sanityCheck() {
	if os.Getenv("DBHOST") == "" {
		panic("DBHOST is undefined")
	}
	if os.Getenv("DBPORT") == "" {
		panic("DBPORT is undefined")
	}
	if os.Getenv("DBUSER") == "" {
		panic("DBUSER is undefined")
	}
	if os.Getenv("DBPASSWORD") == "" { // LRN: TODO: Fatal log also exits the application, this is weird
		log.Println("Warn: DBPASSWORD is undefined")
	}
	if os.Getenv("DBNAME") == "" {
		panic("DBNAME is undefined")
	}
	if os.Getenv("SERVER_ADDRESS") == "" {
		panic("SERVER_ADDRESS is undefined")
	}
	if os.Getenv("SERVER_PORT") == "" {
		panic("SERVER_PORT is undefined")
	}

}

func getClientDb() (*sqlx.DB, error) {
	host := os.Getenv("DBHOST")
	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

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

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	// wiring the application
	dbclient, err := getClientDb() // TODO: how to create connection only when used?
	if err != nil {
		panic(err)
	}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(dbclient))}
	adb := domain.NewAccountRepositoryDb(dbclient)
	account := domain.Account{
		CustomerId:  "2005",
		OpeningDate: "2020-08-09 10:35:22",
		AccountType: "saving",
		Amount:      32432.34,
		Status:      "1",
	}
	adb.Save(account)
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	// starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), router))
}

//SERVER_ADDRESS=localhost SERVER_PORT=8080 DBHOST=localhost DBPORT=5432 DBUSER=postgres DBNAME=banking go run main.go
