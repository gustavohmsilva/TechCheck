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
func (r *Genre) Create(ctx context.Context, g *model.Genre) (*model.Genre, error) {
	// Cria squirrel, etc parar storage
	qry, args, err := squirrel.Insert(
		"Genre",
	).Columns(
		"Name",
	).Values(
		g.Name,
	).ToSql()

	if err != nil {

		return nil, err
	}
	id, err := r.DB.ExecContext(ctx, qry, args...)
	if err != nil {
		return nil, err
	}
	g.ID, err = id.LastInsertId()
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Find ...
func (r *Genre) Find(ctx context.Context, ga *model.GenreArgs) ([]*model.Genre, error) {
	sel := squirrel.Select("*").From("Genre")
	switch {
	case ga.Request.Like != "":
		sel = sel.Where(squirrel.Like{"Name": ga.Request.Like})
	case ga.Request.Offset != 0:
		sel = sel.Offset(ga.Request.Offset)
	case ga.Request.Size != 0:
		sel = sel.Limit(ga.Request.Size)
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
