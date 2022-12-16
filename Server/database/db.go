package Database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// define DB as a pointer to a database object
var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "220402"
	dbname   = "forumdb"
)

// define the init function which connects to the postgresql database
// this will be run before main() as database package will be imported into main package
// then we define DB to point specifically to the forum database (specified by the info in psqlInfo which is sent to sql.Open)
// db.Ping will then attempt to open a connection with the database
// if error occurs, error message will be printed
func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successful Connection")
	}

}
