package api

import (
	"net/http"

	"github.com/gustavohmsilva/TechCheck/model"
	"github.com/gustavohmsilva/TechCheck/rendering"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/gustavohmsilva/TechCheck/util/parser"
	"github.com/labstack/echo/v4"
)

// UserType basically serves as a way to transport tech and other relevante
// references.
type UserType struct {
	userTypeService *tech.UserType
}

// Routes attach in one Genre all related routes
func (ut *UserType) Routes(api *echo.Group) {
	api.POST("", ut.create)
	api.GET("", ut.find)
}

type userTypeCreateRequest struct {
	*tech.UserType
}

type userTypeCreateResponse struct {
	*tech.UserType
}

// create - Insert in the database a valid new User Type
// Success: 200 []model.UserTypeArgs
// Fail: 400 rendering.ResponseError
func (ut *UserType) create(ech echo.Context) error {
	ctx := ech.Request().Context()

	req := new(model.UserTypeArgs)
	err := ech.Bind(&req.UserType)
	if err != nil {
		return ech.JSON(
			http.StatusBadRequest,
			&rendering.ResponseError{
				Errors: []string{err.Error()},
			},
		)
	}

	req.UserType, err = ut.userTypeService.Create(ctx, req)
	if err != nil {
		return ech.JSON(
			http.StatusConflict,
			&rendering.ResponseError{
				Errors: []string{err.Error()},
			},
		)
	}
	return ech.JSON(http.StatusCreated, req.UserType)
}

// find - Search for one or more UserTypes
// Success: 200 []model.UserTypeArgs
// Fail: 400 rendering.ResponseError
// Exception: 500 empty
func (ut *UserType) find(ech echo.Context) error {
	ctx := ech.Request().Context()

	req := new(model.UserTypesArgs)
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
		req.Includes.Count, err = ut.userTypeService.Count(ctx, req)
		if err != nil {
			return ech.JSON(
				http.StatusInternalServerError,
				&rendering.ResponseError{
					Errors: []string{err.Error()},
				},
			)
		}
	}

	req.UserTypes, err = ut.userTypeService.Find(ctx, req)
	if err != nil {
		return ech.NoContent(http.StatusInternalServerError)
	}

	return ech.JSON(http.StatusOK, req)
}
