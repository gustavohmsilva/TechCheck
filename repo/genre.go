package repo

import (
	"context"
	"database/sql"
	"fmt"

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
	if ga.Includes.Like != "" {
		sel = sel.Where(
			squirrel.Like{
				"name": fmt.Sprint(wc, ga.Includes.Like, wc),
			},
		)
	}
	if ga.Includes.Offset != 0 {
		sel = sel.Offset(ga.Includes.Offset)
	}
	if ga.Includes.Size != 0 {
		sel = sel.Limit(ga.Includes.Size)
	}
	qry, args, err := sel.ToSql()
	if err != nil {
		return nil, err
	}
	fmt.Println(qry)
	res, err := r.DB.QueryContext(ctx, qry, args...)
	if err != nil {
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
