package service

import (
	"database/sql"
	"fmt"
)

type Genre struct {
	DB *sql.DB
}

func NewGenre(db *sql.DB) (Genre, error) {
	return Genre{db}, nil
}

// Insert convert the Genre model into a SQL statement and store
// it in the database
func (db *Genre) Insert(stmt string, params []interface{}) (int, error) {
	fmt.Println(stmt, params)
	_, err := db.DB.Exec(stmt, params[0])
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return 1, nil
}
