package utility

import (
	"log"
	"database/sql"
	"time"
)

// Check if connection to database can be established
// It returns error in case the connection cannot be established
func CheckDbConnection(db *sql.DB) (error) {
	var err error

	for dbConnectionTries := 0; dbConnectionTries < 60; dbConnectionTries++ {
		if err = db.Ping(); err == nil {
			log.Println("Connection to database established!")
			return nil // In case of success, we are returning null value as error
		}
		log.Printf("Database not ready, waiting... attempt %d/60", dbConnectionTries + 1)
		time.Sleep(1 * time.Second) // Sleep for 1 second before retrying
	}

	return err
}