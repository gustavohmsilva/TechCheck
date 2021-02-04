package tech

import "context"

// Genre describes a genre or category of book
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type genreRepository interface {
	Create(ctx context.Context, g *Genre) (*Genre, error)
	Find(ctx context.Context, q QueryOptions) ([]*Genre, error)
}

type GenreService struct {
	repo genreRepository
}

func NewGenreService(r genreRepository) *GenreService {
	return &GenreService{r}
}

func (s *GenreService) Create(ctx context.Context, g *Genre) (*Genre, error) {
	// validaçao, cache, etc

	g, err := s.repo.Create(ctx, g)

	// Alguma outra coisa

	return g, err
}

func (s *GenreService) Find(ctx context.Context, q QueryOptions) ([]*Genre, error) {
	g, err := s.repo.Find(ctx, q)

	return g, err
}
