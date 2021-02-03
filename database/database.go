package database

import (
	"database/sql"
	"fmt"

	// MariaDB Driver
	_ "github.com/go-sql-driver/mysql"
)

// Database holds the connection and necessary data for connecting to the
// database
type Database struct {
	DB *sql.DB
}

// NewDatabase instanciate a new Database object
func NewDatabase() (*Database, error) {
	db, err := sql.Open(
		"mysql",
		"root:root@/techcheck",
	)
	if err != nil {
		fmt.Println("Connection to PostgreSQL database failed!")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to PostgreSQL database failed!")
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL database")
	return &Database{db}, nil
}
