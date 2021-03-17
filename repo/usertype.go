package repo

import (
	"context"
	"database/sql"

	"github.com/gustavohmsilva/TechCheck/model"
)

// UserType vai persistir os genres
type UserType struct {
	DB *sql.DB
}

// NewUserType ...
func NewUserType(db *sql.DB) *UserType {
	return &UserType{db}
}

func (ut *UserType) Find(
	ctx context.Context, utsa *model.UserTypesArgs,
) (
	[]*model.UserType, error,
) {
	/*t := reflect.TypeOf(ga.ID)
	f := t.Field(1) // Name
	sel := squirrel.Select(all).From(genre)
	if ga.Request.Like != "" {
		sel = sel.Where(
			squirrel.Like{
				f.Name: fmt.Sprint(wc, ga.Request.Like, wc),
			},
		)
	}
	if ga.Request.Offset != 0 {
		sel = sel.Offset(ga.Request.Offset)
	}
	if ga.Request.Size != 0 {
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
	}*/
	return nil, nil
}
