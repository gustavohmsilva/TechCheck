package tech

import (
	"context"

	"github.com/gustavohmsilva/TechCheck/model"
)

type BookRepository interface {
	Create(ctx context.Context, b *model.Book) (*model.Book, error)
	Find(ctx context.Context) ([]*model.Book, error)
}

type Book struct {
	repo BookRepository
}

func NewBook(r BookRepository) *Book {
	return &Book{r}
}

func (b *Book) Create(ctx context.Context, g *model.Book) (*model.Book, error) {
	// validaçao, cache, etc

	g, err := b.repo.Create(ctx, g)

	// Alguma outra coisa

	return g, err
}

func (b *Book) Find(ctx context.Context) ([]*model.Book, error) {
	g, err := b.repo.Find(ctx)

	return g, err
}
