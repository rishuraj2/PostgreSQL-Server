package utility

import (
	"database/sql"
)

// Check if the table exists
// Returns true in case the table is found and vice-versa
func CheckTableExists(db *sql.DB, schemaType string, tableName string) bool {

	exists := false

	query := `SELECT EXISTS (
	SELECT FROM information_schema.tables
	WHERE table_schema = $1 AND table_name = $2
	)`

	db.QueryRow(query, schemaType, tableName).Scan(&exists)
	
	return exists;
}