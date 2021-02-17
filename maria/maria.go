// Package maria just creates a db connection to a mariaDB database
package maria

import (
	"database/sql"
	"fmt"
	"log"

	// Database driver
	_ "github.com/go-sql-driver/mysql"
)

// NewDB returns a testes connection to the database
func NewDB() (*sql.DB, error) {
	connLine := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase,
	)
	db, err := sql.Open("mysql", connLine)
	if err != nil {
		log.Println("Connection to MariaDB database failed!")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Connection to MariaDB database failed!")
		return nil, err
	}
	return db, nil
}
