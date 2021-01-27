package api

import (
	"net/http"
	"strconv"

	"github.com/gustavohmsilva/TechCheck/models"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

// Genre basically serves as a way to transport tech and other relevante
// references.
type Genre struct {
	tech  *tech.Tech
	group *echo.Group
	RequestParameters
}

// RequestParameters contains basic parameters for when querying data
type RequestParameters struct {
	Find   string
	Amount int
	Offset int
	Valid  []error
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
	genreModel := new(models.Genre)
	err := cnx.Bind(&genreModel)
	if err != nil {
		return err
	}
	if genreModel.Name == "" {
		return cnx.JSON(
			http.StatusBadRequest,
			&models.ResponseError{Err: "No genre provided"},
		)
	}
	if genreModel.ID != 0 {
		genreModel.ID = 0
	}
	return cnx.JSON(http.StatusOK, genreModel)
}

// retrieveGenre will retrieve one or more genres from the database depending
// on the parameter used for search.
// TODO: Remember to implement a limit
func (genre *Genre) retrieveGenre(cnx echo.Context) error {
	genre.Find = cnx.QueryParam("find")
	var err error
	genre.Amount, err = strconv.Atoi(cnx.QueryParam("amount"))
	if err != nil {
		genre.Valid = append(genre.Valid, err)
	}
	genre.Offset, err = strconv.Atoi(cnx.QueryParam("offset"))
	if err != nil {
		genre.Valid = append(genre.Valid, err)
	}
	return cnx.JSON(http.StatusOK, genre)
}
