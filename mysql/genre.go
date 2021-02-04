package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/gustavohmsilva/TechCheck/tech"
)

// GenreRepository vai persistir os genres
type GenreRepository struct {
	DB *sql.DB
}

func NewGenreRepository(db *sql.DB) *GenreRepository {
	return &GenreRepository{db}
}

// Create  ...
func (r *GenreRepository) Create(ctx context.Context, g *tech.Genre) (*tech.Genre, error) {
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
func (r *GenreRepository) Find(ctx context.Context, q tech.QueryOptions) ([]*tech.Genre, error) {
	// etc etc, limit, query, etc
	qry, args, err := squirrel.Select("genre").ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.DB.QueryContext(ctx, qry, args...)

	if err != nil {
		return nil, err
	}

	return []*tech.Genre{}, errors.New("not implemented")
}
