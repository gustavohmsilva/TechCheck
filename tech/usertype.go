package tech

import (
	"context"
	"errors"

	"github.com/gustavohmsilva/TechCheck/model"
)

type UserTypeRepository interface {
	Create(
		ctx context.Context, uta *model.UserTypeArgs,
	) (
		*model.UserType, error,
	)
	Find(
		ctx context.Context, utsa *model.UserTypesArgs,
	) (
		[]*model.UserType, error,
	)
	Count(ctx context.Context, utca *model.UserTypesArgs) (uint64, error)
}

type UserType struct {
	repo UserTypeRepository
}

func NewUserType(r UserTypeRepository) *UserType {
	return &UserType{r}
}

func (ut *UserType) Create(
	ctx context.Context, uta *model.UserTypeArgs,
) (
	*model.UserType, error,
) {
	if uta.UserType.ID != 0 {
		uta.UserType.ID = 0
	}

	if uta.UserType.Name == "" {
		return nil, errors.New("no user type name provided")
	}

	storedUserType, err := ut.repo.Create(ctx, uta)
	if err != nil {
		return nil, err
	}

	return storedUserType, nil
}

func (ut *UserType) Find(
	ctx context.Context, utsa *model.UserTypesArgs,
) (
	[]*model.UserType, error,
) {
	if utsa.Includes.Size > 50 || utsa.Includes.Size < 1 {
		utsa.Includes.Size = 50
	}

	foundUserTypes, err := ut.repo.Find(ctx, utsa)
	if err != nil {
		return nil, err
	}

	return foundUserTypes, nil
}

func (ut *UserType) Count(
	ctx context.Context, utca *model.UserTypesArgs,
) (
	uint64, error,
) {
	if len(utca.Includes.Like) < 3 {
		return 0, errors.New(
			"count require at least three characters in field LIKE",
		)
	}

	return ut.repo.Count(ctx, utca)
}
