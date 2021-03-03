package tech

import (
	"context"
	"errors"
	"fmt"

	"github.com/gustavohmsilva/TechCheck/model"
)

type genreRepository interface {
	Create(ctx context.Context, g *model.Genre) (*model.Genre, error)
	Find(ctx context.Context, ga *model.GenreArgs) ([]*model.Genre, error)
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
	ctx context.Context, g *model.Genre,
) (
	*model.Genre, error,
) {
	if g.ID != 0 {
		g.ID = 0
	}
	if g.Name == "" {
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
	ctx context.Context, ga *model.GenreArgs,
) (
	[]*model.Genre, error,
) {
	if len(ga.Request.Like) > 128 {
		return nil, errors.New("Search string too big")
	}
	if ga.Request.Size > 50 {
		ga.Request.Size = 50
	}
	gs, err := s.repo.Find(ctx, ga)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return gs, nil
}
