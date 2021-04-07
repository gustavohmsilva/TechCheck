package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/gustavohmsilva/TechCheck/model"
	l "github.com/sirupsen/logrus"
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

func (r *Book) Count(
	ctx context.Context,
	bca *model.BooksArgs,
) (
	uint64,
	error,
) {
	ct := squirrel.Select(
		"count(Id)",
	).From(
		"Book",
	).Where(
		// TODO: Improve this where to take into account isbn and author
		// as well. Currently only count in database using name.
		squirrel.Like{
			"Name": fmt.Sprint(
				wc,
				bca.Includes.Like,
				wc,
			),
		},
	)

	qry, args, err := ct.ToSql()
	if err != nil {
		l.Errorf("QUERY ERROR: %s", err.Error())
		return 0, err
	}

	l.Infof("QUERY: %s", qry)
	res := r.DB.QueryRowContext(ctx, qry, args...)
	var count uint64
	err = res.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
