package main

import (
	"github.com/gustavohmsilva/TechCheck/api"
	"github.com/gustavohmsilva/TechCheck/maria"
	"github.com/gustavohmsilva/TechCheck/repo"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

func main() {

	// Does the connection to the database, it will panic if impossible to
	// connect.
	database, err := maria.NewDB()
	if err != nil {
		panic(err)
	}

	// Creates the repositories. They contains the functions that work
	// directly to the databases.
	genreRepository := repo.NewGenre(database)
	bookRepository := repo.NewBook(database)

	// Creates the "services" (which the package receive the name of the
	// app itself). It does validation and business logic between the
	// request and the response.
	genreService := tech.NewGenre(genreRepository)
	bookService := tech.NewBook(bookRepository)

	a := &api.API{
		Genres: genreService,
		Books:  bookService,
		// Users...
	}

	e := echo.New()

	a.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
