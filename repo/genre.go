package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/gustavohmsilva/TechCheck/model"
)

// Genre vai persistir os genres
type Genre struct {
	DB *sql.DB
}

// NewGenre ...
func NewGenre(db *sql.DB) *Genre {
	return &Genre{db}
}

// Create  ...
func (r *Genre) Create(ctx context.Context, g *model.Genre) (*model.Genre, error) {
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
func (r *Genre) Find(ctx context.Context) ([]*model.Genre, error) {
	// etc etc, limit, query, etc
	qry, args, err := squirrel.Select("genre").ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.DB.QueryContext(ctx, qry, args...)

	if err != nil {
		return nil, err
	}

	return []*model.Genre{}, errors.New("not implemented")
}
