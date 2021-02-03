package api

import (
	"net/http"
	"strconv"

	"github.com/gustavohmsilva/TechCheck/model"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

// Genre basically serves as a way to transport tech and other relevante
// references.
type Genre struct {
	tech  *tech.Tech
	group *echo.Group
	requestParameters
}

// RequestParameters contains basic parameters for when querying data
type requestParameters struct {
	find   string
	amount int
	offset int
	valid  []error
}

// NewGenre instanciate a new genre
func NewGenre(t *tech.Tech, e *echo.Echo) *Genre {
	return &Genre{tech: t, group: e.Group("/genre")}
}

// Routes attach in one Genre all related routes
func (genre *Genre) Routes() {
	genre.group.POST("", genre.createGenre)
	genre.group.GET("", genre.retrieveGenre)
}

// createGenre retrieves the JSON sent by the user with a new genre and send it
// to be stored in the database.
func (genre *Genre) createGenre(cnx echo.Context) error {
	genreModel := new(model.Genre)
	err := cnx.Bind(&genreModel)
	if err != nil {
		return err
	}
	if genreModel.Name == "" {
		return cnx.JSON(
			http.StatusBadRequest,
			&model.ResponseError{Err: "No genre provided"},
		)
	}

	md, err := model.NewDbMap(genre.tech.Database.DB)
	dbg, err := md.Create(genreModel)
	if err != nil {
		return err
	}
	// TODO: Código para model
	return cnx.JSON(http.StatusOK, dbg)
}

// retrieveGenre will retrieve one or more genres from the database depending
// on the parameter used for search.
// TODO: Remember to implement a limit
func (genre *Genre) retrieveGenre(cnx echo.Context) error {
	genre.find = cnx.QueryParam("find")
	var err error
	genre.amount, err = strconv.Atoi(cnx.QueryParam("amount"))
	if err != nil {
		genre.valid = append(genre.valid, err)
	}
	genre.offset, err = strconv.Atoi(cnx.QueryParam("offset"))
	if err != nil {
		genre.valid = append(genre.valid, err)
	}
	return cnx.JSON(http.StatusOK, genre)
}
