package api

import (
	"net/http"

	"github.com/gustavohmsilva/TechCheck/model"
	"github.com/gustavohmsilva/TechCheck/rendering"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

// Genre basically serves as a way to transport tech and other relevante
// references.
type Genre struct {
	genreService *tech.Genre
}

// Routes attach in one Genre all related routes
func (g *Genre) Routes(api *echo.Group) {
	api.POST("", g.create)
	api.GET("", g.find)
}

// "models" de camada de transporte, aqui dependendo de crud ou alguma layer
// extra de transporte, poderia usar genres diretamente ou usar para cada chamada
// uma struct
type genreCreateRequest struct {
	*tech.Genre

	// outras opts para passar no request
	// ...
}

type genreCreateResponse struct {
	*tech.Genre
}

// createGenre retrieves the JSON sent by the user with a new genre and send it
// to be stored in the database.
func (g *Genre) create(ech echo.Context) error {
	ctx := ech.Request().Context()
	// poderia ser o tech.Genre
	req := new(model.Genre)
	err := ech.Bind(&req)
	if err != nil {
		return err
	}

	if req.Name == "" {
		return ech.JSON(
			http.StatusBadRequest,
			"No genre provided",
		)
	}

	createdGenre, err := g.genreService.Create(ctx, req)
	if err != nil {
		return ech.JSON(
			http.StatusConflict,
			&rendering.ResponseError{
				Err: err.Error(),
			},
		)
	}
	return ech.JSON(http.StatusOK, createdGenre)
}

// find will retrieve one or more genres from the database depending
// on the parameter used for search.
// TODO: Remember to implement a limit
func (g *Genre) find(ech echo.Context) error {
	ctx := ech.Request().Context()
	req := new(model.GenreArgs)
	err := ech.Bind(&req)
	if err != nil {
		return err
	}
	genres, err := g.genreService.Find(ctx, req)
	if err != nil {
		return err
	}
	return ech.JSON(http.StatusOK, genres)
}
