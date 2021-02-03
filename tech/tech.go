package tech

import (
	"github.com/gustavohmsilva/TechCheck/database"
)

// Tech is an istance of the tech application
type Tech struct {
	Database *database.Database
}

// NewTech instanciate a new Tech object
func NewTech(db *database.Database) (*Tech, error) {
	var T Tech
	var err error
	T.Database = db
	return &T, err
}
