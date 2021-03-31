package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/gustavohmsilva/TechCheck/model"
	l "github.com/sirupsen/logrus"
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
func (r *Genre) Create(
	ctx context.Context, g *model.GenreArgs,
) (
	*model.Genre, error,
) {
	// Cria squirrel, etc parar storage
	qry, args, err := squirrel.Insert(
		"Genre",
	).Columns(
		"Name",
	).Values(
		g.Genre.Name,
	).ToSql()

	if err != nil {

		return nil, err
	}
	id, err := r.DB.ExecContext(ctx, qry, args...)
	if err != nil {
		return nil, err
	}
	g.Genre.ID, err = id.LastInsertId()
	if err != nil {
		return nil, err
	}
	return g.Genre, nil
}

// Find ...
func (r *Genre) Find(
	ctx context.Context,
	ga *model.GenresArgs,
) (
	[]*model.Genre,
	error,
) {
	sel := squirrel.Select(all).From(genre)
	if ga.Includes.ID != 0 {
		sel = sel.Where(squirrel.Eq{"Id": ga.Includes.ID})
	} else {
		if ga.Includes.Like != "" {
			sel = sel.Where(
				squirrel.Like{
					"name": fmt.Sprint(
						wc,
						ga.Includes.Like,
						wc,
					),
				},
			)
		}

		if ga.Includes.Offset != 0 {
			sel = sel.Offset(ga.Includes.Offset)
		}

		sel = sel.Limit(ga.Includes.Size)
	}

	qry, args, err := sel.ToSql()
	if err != nil {
		l.Errorf("QUERY ERROR: %s", err.Error())
		return nil, err
	}

	l.Infof("QUERY: %s", qry)
	res, err := r.DB.QueryContext(ctx, qry, args...)
	if err != nil {
		l.Errorf("QUERY ERROR: %s", err.Error())
		return nil, err
	}

	gs := make([]*model.Genre, 0)
	for res.Next() {
		var g model.Genre
		err := res.Scan(&g.ID, &g.Name)
		if err != nil {
			continue
		}
		gs = append(gs, &g)
	}
	return gs, nil
}

func (r *Genre) Count(
	ctx context.Context, utca *model.GenresArgs,
) (
	uint64, error,
) {
	ct := squirrel.Select(
		"count(Id)",
	).From(
		"Genre",
	).Where(
		squirrel.Like{
			"Name": fmt.Sprint(
				wc,
				utca.Includes.Like,
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
