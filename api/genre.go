package api

import (
	"net/http"

	"github.com/gustavohmsilva/TechCheck/model"
	"github.com/gustavohmsilva/TechCheck/rendering"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/gustavohmsilva/TechCheck/util/parser"
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

// create - Insert in the database a valid new genre
// Success: 200 []model.Genre
// Fail: 400 rendering.ResponseError
func (g *Genre) create(ech echo.Context) error {
	ctx := ech.Request().Context()
	// poderia ser o tech.Genre
	req := new(model.GenreArgs)
	err := ech.Bind(&req.Genre)
	if err != nil {
		return ech.JSON(
			http.StatusBadRequest,
			&rendering.ResponseError{
				Err: err.Error(),
			},
		)
	}

	req.Genre, err = g.genreService.Create(ctx, req)
	if err != nil {
		return ech.JSON(
			http.StatusConflict,
			&rendering.ResponseError{
				Err: err.Error(),
			},
		)
	}
	return ech.JSON(http.StatusOK, req.Genre)
}

// find - Search for one or more genres
// Success: 200 []model.GenreArgs
// Fail: 400 rendering.ResponseError
// Exception: 500 empty
func (g *Genre) find(ech echo.Context) error {
	ctx := ech.Request().Context()

	req := new(model.GenresArgs)

	req.Includes.Like = ech.QueryParam("like")

	var re rendering.ResponseError
	req.Includes.Size, re = parser.Uint64(ech.QueryParam("size"))
	if (re != rendering.ResponseError{}) {
		return ech.JSON(http.StatusBadRequest, re)
	}

	req.Includes.Offset, re = parser.Uint64(ech.QueryParam("offset"))
	if (re != rendering.ResponseError{}) {
		return ech.JSON(http.StatusBadRequest, re)
	}

	if req.Includes.Offset < 1 {
		var err error
		req.Includes.Count, err = g.genreService.Count(ctx, req)
		if err != nil {
			return ech.JSON(
				http.StatusInternalServerError,
				&rendering.ResponseError{
					Err: err.Error(),
				},
			)
		}
	}

	var err error
	req.Genres, err = g.genreService.Find(ctx, req)
	if err != nil {
		return ech.NoContent(http.StatusInternalServerError)
	}

	return ech.JSON(http.StatusOK, req)
}
