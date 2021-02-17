package maria

import "os"

var (
	dbHost     = os.Getenv("DB_HOST")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbPort     = os.Getenv("DB_PORT")
)
