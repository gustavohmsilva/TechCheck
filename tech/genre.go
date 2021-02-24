package tech

import (
	"context"
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

func (s *Genre) Create(ctx context.Context, g *model.Genre) (*model.Genre, error) {
	// validaçao, cache, etc

	storedGenre, err := s.repo.Create(ctx, g)
	if err != nil {
		return nil, err
	}
	// Alguma outra coisa

	return storedGenre, nil
}

func (s *Genre) Find(ctx context.Context, ga *model.GenreArgs) ([]*model.Genre, error) {
	gs, err := s.repo.Find(ctx, ga)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return gs, nil
}
