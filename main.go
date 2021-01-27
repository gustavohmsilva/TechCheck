package main

import (
	"github.com/gustavohmsilva/TechCheck/api"
	"github.com/gustavohmsilva/TechCheck/database"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := database.NewDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}
	t, err := tech.NewTech(db)
	if err != nil {
		e.Logger.Fatal(err)
	}
	genreGroup := api.NewGenre(t, e)
	genreGroup.Routes()
	e.Logger.Fatal(e.Start(":8080"))
}
