package database

import "database/sql"

// Database holds the connection and necessary data for connecting to the
// database
type Database struct {
	Db *sql.DB
}

// NewDatabase instanciate a new Database object
func NewDatabase() (*Database, error) {
	// TODO: Create Database Connection Code
	return &Database{nil}, nil
}
