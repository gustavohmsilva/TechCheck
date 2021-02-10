// Package maria just creates a db connection to a mariaDB database
package maria

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// NewDB returns a testes connection to the database
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/techcheck")
	if err != nil {
		log.Fatal("Connection to MariaDB database failed!")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to MariaDB database failed!")
	}
	return db
}
