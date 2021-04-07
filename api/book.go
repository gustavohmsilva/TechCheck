package api

import (
	"net/http"

	"github.com/gustavohmsilva/TechCheck/model"
	"github.com/gustavohmsilva/TechCheck/rendering"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/gustavohmsilva/TechCheck/util/parser"
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
func (b *Book) find(ech echo.Context) error {
	ctx := ech.Request().Context()
	req := new(model.BooksArgs)

	req.Includes.Name = ech.QueryParam("name")
	req.Includes.Name = ech.QueryParam("author")
	req.Includes.Like = ech.QueryParam("like")

	var err error
	requestFailed := new(rendering.ResponseError)
	req.Includes.ID, err = parser.Uint64(
		ech.QueryParam("id"),
		"id",
	)
	if err != nil {
		requestFailed.Errors = append(requestFailed.Errors, err.Error())
	}

	req.Includes.ISBN, err = parser.Uint64(
		ech.QueryParam("isbn"),
		"isbn",
	)
	if err != nil {
		requestFailed.Errors = append(requestFailed.Errors, err.Error())
	}

	req.Includes.Size, err = parser.Uint64(
		ech.QueryParam("size"),
		"size",
	)
	if err != nil {
		requestFailed.Errors = append(requestFailed.Errors, err.Error())
	}

	req.Includes.Offset, err = parser.Uint64(
		ech.QueryParam("offset"),
		"offset",
	)
	if err != nil {
		requestFailed.Errors = append(requestFailed.Errors, err.Error())
	}

	if len(requestFailed.Errors) > 0 {
		return ech.JSON(http.StatusBadRequest, requestFailed)
	}

	if req.Includes.Offset < 1 {
		var err error
		req.Includes.Count, err = b.bookService.Count(ctx, req)
		if err != nil {
			return ech.JSON(
				http.StatusInternalServerError,
				&rendering.ResponseError{
					Errors: []string{
						err.Error(),
					},
				},
			)
		}
	}
	// TODO: implementation of bookservice.find has to be entirely refactored
	//req.Books, err = b.bookService.Find(ctx, req)
	if err != nil {
		return ech.NoContent(http.StatusInternalServerError)
	}

	return ech.JSON(http.StatusOK, req)
}
