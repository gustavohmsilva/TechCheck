package api

import (
	"net/http"

	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

// UserType basically serves as a way to transport tech and other relevante
// references.
type UserType struct {
	usertTypeService *tech.UserType
}

// Routes attach in one Genre all related routes
func (ut *UserType) Routes(api *echo.Group) {
	api.GET("", ut.find)
}

type userTypeCreateRequest struct {
	*tech.UserType
}

type userTypeCreateResponse struct {
	*tech.UserType
}

func (ut *UserType) find(ech echo.Context) error {

	return ech.NoContent(http.StatusOK)
}
