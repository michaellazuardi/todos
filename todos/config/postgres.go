package config

import (
	"database/sql"
	"fmt"
	"os"
)

//we need to create a function to fetch the database from the postgreSQL
func getPostgresDB() (*sql.DB, error) {
	//the contents of the Getenv MUST BY DEFAULT be inserted as an argument of Getenv
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	//why do we use Sprintf instead of Printf??

	//the sprintf argument MUST BE INSERTED BY DEFAULT
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	db, err := createConnection(desc)

	return db, err

}

func createConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)

	}
	return db, nil
}
