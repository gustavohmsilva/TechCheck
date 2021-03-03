// Package api ...
// - nao sabe o que é SQL nem cache nem persistencia nem DB
// - sabe validar requests http
// - Sabe transformar os requests e responses para coisas da package tech
// - e fazer chamadas na package tech para servir dados da package tech
package api

import (
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

type API struct {
	Genres    *tech.Genre
	Books     *tech.Book
	UserTypes *tech.UserType
}

func (a *API) Routes(e *echo.Echo) error {
	g := &Genre{genreService: a.Genres}
	g.Routes(e.Group("/api/v1/genre"))
	b := &Book{bookService: a.Books}
	b.Routes(e.Group("/api/v1/book"))
	ut := &UserType{usertTypeService: a.UserTypes}
	ut.Routes(e.Group("/api/v1/usertype"))

	return nil
}
