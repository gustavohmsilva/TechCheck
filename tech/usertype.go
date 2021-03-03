package tech

import (
	"context"

	"github.com/gustavohmsilva/TechCheck/model"
)

type UserTypeRepository interface {
	Find(ctx context.Context, args *model.UserTypeArgs) ([]*model.UserType, error)
}

type UserType struct {
	repo UserTypeRepository
}

func NewUserType(r UserTypeRepository) *UserType {
	return &UserType{r}
}

func (ut *UserType) Find(ctx context.Context, args *model.UserTypeArgs) ([]*model.UserType, error) {
	var ret []*model.UserType
	return ret, nil
}
