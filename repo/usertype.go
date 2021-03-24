package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/gustavohmsilva/TechCheck/model"
	l "github.com/sirupsen/logrus"
)

// UserType vai persistir os genres
type UserType struct {
	DB *sql.DB
}

// NewUserType ...
func NewUserType(db *sql.DB) *UserType {
	return &UserType{db}
}

func (ut *UserType) Create(
	ctx context.Context, uta *model.UserTypeArgs,
) (
	*model.UserType, error,
) {
	qry, args, err := squirrel.Insert(
		"UserType",
	).Columns(
		"Name",
	).Values(
		uta.UserType.Name,
	).ToSql()

	if err != nil {
		return nil, err
	}
	id, err := ut.DB.ExecContext(ctx, qry, args...)
	if err != nil {
		return nil, err
	}
	uta.UserType.ID, err = id.LastInsertId()
	if err != nil {
		return nil, err
	}
	return uta.UserType, nil
}

func (ut *UserType) Find(
	ctx context.Context, utsa *model.UserTypesArgs,
) (
	[]*model.UserType, error,
) {
	sel := squirrel.Select(all).From("UserType")
	if utsa.Includes.ID != 0 {
		sel = sel.Where(squirrel.Eq{"Id": utsa.Includes.ID})
	} else {
		if utsa.Includes.Like != "" {
			sel = sel.Where(
				squirrel.Like{
					"Name": fmt.Sprint(
						wc,
						utsa.Includes.Like,
						wc,
					),
				},
			)
		}
		if utsa.Includes.Offset != 0 {
			sel = sel.Offset(utsa.Includes.Offset)

		}
		sel = sel.Limit(utsa.Includes.Size)
	}

	qry, args, err := sel.ToSql()
	if err != nil {
		l.Errorf("QUERY ERROR: %s", err.Error())
		return nil, err
	}

	l.Infof("QUERY: %s", qry)
	res, err := ut.DB.QueryContext(ctx, qry, args...)
	if err != nil {
		l.Errorf("QUERY ERROR: %s", err.Error())
		return nil, err
	}
	retUserType := make([]*model.UserType, 0)
	for res.Next() {
		var newUserType model.UserType
		err := res.Scan(&newUserType.ID, &newUserType.Name)
		if err != nil {
			continue
		}
		retUserType = append(retUserType, &newUserType)
	}
	return retUserType, nil
}
