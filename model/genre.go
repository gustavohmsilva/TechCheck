package model

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/gustavohmsilva/TechCheck/service"
)

// DbMap ...
type DbMap struct {
	Db *sql.DB
}

// NewDbMap ...
func NewDbMap(db *sql.DB) (*DbMap, error) {
	return &DbMap{Db: db}, nil
}

// Genre describes a genre or category of book
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Create will produce
func (db *DbMap) Create(g *Genre) (*Genre, error) {
	// validações adicionais
	stmt, arg, err := sq.Insert(
		"Genre",
	).Columns(
		"Name",
	).Values(
		g.Name,
	).ToSql()
	genreSvc, err := service.NewGenre(db.Db)
	g.ID, err = genreSvc.Insert(stmt, arg)
	if err != nil {
		return &Genre{}, err
	}
	return g, nil
}
