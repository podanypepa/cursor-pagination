package db

import (
	"database/sql"
	"fmt"
)

var (
	dbhost     = ""
	dbport     = 5432
	dbuser     = ""
	dbpassword = ""
	dbname     = ""

	DB *sql.DB
)

// Connect to the PostgreSQL
func Connect() (err error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, dbname,
	)
	if DB, err = sql.Open("postgres", psqlInfo); err != nil {
		return err
	}
	return nil
}
