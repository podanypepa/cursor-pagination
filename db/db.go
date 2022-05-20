package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbhost     = ""
	dbport     = ""
	dbuser     = ""
	dbpassword = ""
	dbname     = ""

	DB *sql.DB
)

const ENVFILE = ".env"

func init() {
	if _, err := os.Stat(ENVFILE); os.IsExist(err) {
		if err := godotenv.Load(ENVFILE); err != nil {
			log.Fatal(err)
		}
	}
	dbhost = os.Getenv("DBHOST")
	dbport = os.Getenv("DBPORT")
	dbuser = os.Getenv("DBUSER")
	dbpassword = os.Getenv("DBPASSWORD")
	dbname = os.Getenv("DBNAME")
}

// Connect to the PostgreSQL
func Connect() (err error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, dbname,
	)
	if DB, err = sql.Open("postgres", psqlInfo); err != nil {
		return err
	}
	return nil
}
