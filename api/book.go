package api

import (
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

// Book basically serves as a way to transport tech and other relevante
// references.
type Book struct {
	bookService *tech.Book
}

// Routes attach in one Genre all related routes
func (b *Book) Routes(api *echo.Group) {
	api.POST("", b.create)
	api.GET("", b.find)
}

// Create is the controler that deals with POST for Book
func (b *Book) create(cnx echo.Context) error {
	return nil
}

// Create is the controler that deals with POST for Book
func (b *Book) find(cnx echo.Context) error {
	return nil
}
