package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/gustavohmsilva/TechCheck/model"
)

// Book vai persistir os genres
type Book struct {
	DB *sql.DB
}

// NewBook ...
func NewBook(db *sql.DB) *Book {
	return &Book{db}
}

// Create  ...
func (r *Book) Create(ctx context.Context, g *model.Book) (*model.Book, error) {
	// Cria squirrel, etc parar storage
	qry, args, err := squirrel.Insert("genre").ToSql()
	if err != nil {
		return nil, err
	}
	// Insert
	_, err = r.DB.ExecContext(ctx, qry, args...)

	return g, err
}

// Find ...
func (r *Book) Find(ctx context.Context) ([]*model.Book, error) {
	// etc etc, limit, query, etc
	qry, args, err := squirrel.Select("genre").ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.DB.QueryContext(ctx, qry, args...)

	if err != nil {
		return nil, err
	}

	return []*model.Book{}, errors.New("not implemented")
}
