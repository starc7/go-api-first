package connectionDB

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "12345"
    dbname   = "postgres"
)

func GetDB() (*sql.DB, error) {
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	var err error

	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected with program")
	return db, nil
}