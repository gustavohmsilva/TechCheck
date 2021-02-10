package tech

import (
	"context"

	"github.com/gustavohmsilva/TechCheck/model"
)

type genreRepository interface {
	Create(ctx context.Context, g *model.Genre) (*model.Genre, error)
	Find(ctx context.Context) ([]*model.Genre, error)
}

type Genre struct {
	repo genreRepository
}

func NewGenre(r genreRepository) *Genre {
	return &Genre{r}
}

func (s *Genre) Create(ctx context.Context, g *model.Genre) (*model.Genre, error) {
	// validaçao, cache, etc

	g, err := s.repo.Create(ctx, g)

	// Alguma outra coisa

	return g, err
}

func (s *Genre) Find(ctx context.Context) ([]*model.Genre, error) {
	g, err := s.repo.Find(ctx)

	return g, err
}
