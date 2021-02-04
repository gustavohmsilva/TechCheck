package api

import (
	"net/http"
	"strconv"

	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

// Genre basically serves as a way to transport tech and other relevante
// references.
type Genre struct {
	genreService *tech.GenreService
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
func (g *Genre) create(cnx echo.Context) error {
	// poderia ser o tech.Genre
	req := new(genreCreateRequest)
	err := cnx.Bind(req)
	if err != nil {
		return err
	}

	// Esse tipo de validaçao poderia ficar aqui mas pessoalmente gosto de
	// fazer no domain porque se tiver outra camada de transporte evita ter que
	// repetir certas verificaçoes
	if req.Name == "" {
		return cnx.JSON(
			http.StatusBadRequest,
			"No genre provided",
		)
	}

	createdGenre, err := g.genreService.Create(cnx.Request().Context(), req.Genre)
	if err != nil {
		return err
	}

	return cnx.JSON(http.StatusOK, createdGenre)
}

// find will retrieve one or more genres from the database depending
// on the parameter used for search.
// TODO: Remember to implement a limit
func (g *Genre) find(cnx echo.Context) error {
	q := tech.QueryOptions{}
	q.Find = cnx.QueryParam("find")
	var err error

	var errs []error
	q.Amount, err = strconv.Atoi(cnx.QueryParam("amount"))
	if err != nil {
		errs = append(errs, err)
	}
	q.Offset, err = strconv.Atoi(cnx.QueryParam("offset"))
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		// return errors Baseado no slice errs
	}
	genres, err := g.genreService.Find(cnx.Request().Context(), q)
	if err != nil {
		return err
	}

	return cnx.JSON(http.StatusOK, genres)
}
