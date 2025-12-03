package utility

import (
	"database/sql"
)

var (
	dbHost = GetEnv("DB_HOST", "localhost")
	dbPort = GetEnv("DB_PORT", "5432")
	dbName = GetEnv("DB_NAME", "postgres_db")
	dbUser = GetEnv("DB_USER", "postgres_user")
	dbPassword = GetEnv("DB_PASSWORD", "postgres_password")
	sslmode = "disable"
)

// This function helps in connecting to database.
// It returns the database if connected, otherwise will return error.
func ConnectDB() (*sql.DB, error) {

	// connection string looks like "postgres://postgres_user:postgres_password@localhost:5432/postgres_db?sslmode=disables"
	connectionString := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + sslmode

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	return db, nil
}