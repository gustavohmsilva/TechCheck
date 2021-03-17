package tech

import (
	"context"
	"errors"
	"fmt"

	"github.com/gustavohmsilva/TechCheck/model"
)

type genreRepository interface {
	Create(ctx context.Context, g *model.GenreArgs) (*model.Genre, error)
	Find(ctx context.Context, ga *model.GenresArgs) ([]*model.Genre, error)
}

type Genre struct {
	repo genreRepository
}

func NewGenre(r genreRepository) *Genre {
	return &Genre{r}
}

// Create validate the terms received by the controller and if they are usable,
// try to create such genre in the database.
func (s *Genre) Create(
	ctx context.Context, g *model.GenreArgs,
) (
	*model.Genre, error,
) {
	if g.Genre.ID != 0 {
		g.Genre.ID = 0
	}
	if g.Genre.Name == "" {
		return nil, errors.New("No genre provided")
	}
	storedGenre, err := s.repo.Create(ctx, g)
	if err != nil {
		return nil, err
	}
	return storedGenre, nil
}

// Find validate the terms received by the controller and if they are usable,
// performs a database search for a set of genres.
func (s *Genre) Find(
	ctx context.Context, ga *model.GenresArgs,
) (
	[]*model.Genre, error,
) {
	if len(ga.Includes.Like) > 128 {
		return nil, errors.New("Search string too big")
	}
	if ga.Includes.Size > 50 {
		ga.Includes.Size = 50
	}
	gs, err := s.repo.Find(ctx, ga)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return gs, nil
}
