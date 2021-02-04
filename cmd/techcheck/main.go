package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gustavohmsilva/TechCheck/api"
	"github.com/gustavohmsilva/TechCheck/mysql"
	"github.com/gustavohmsilva/TechCheck/tech"
	"github.com/labstack/echo/v4"
)

func main() {
	// a DSN vai vir de uma env e poderia ser inicializado em outra package
	// em alguns casos vi uma package config com coisas tipo
	// db, err := config.InitDB()
	db, err := sql.Open("mysql", "root:root@techcheck")
	if err != nil {
		log.Fatal("Connection to PostgreSQL database failed!")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to PostgreSQL database failed!")
	}

	// Inicia cada um dos "serviços" injectando uma mysql.repository..
	// ou poderia injetar outras coisas cache, outro tipo de repositorio etc
	genresSvc := tech.NewGenreService(mysql.NewGenreRepository(db))

	a := &api.API{
		Genres: genresSvc,
		// Books:
		// Users:
	}

	// Se tivesse uma outra camada de transporte tipo grpc, daria para passar
	// tb o mesmo serviço
	//
	// g := &grpc.Coisa{
	//		Genres: genresSvc,
	// }
	// go grpcServe.Start(g)..

	//
	e := echo.New()

	a.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
