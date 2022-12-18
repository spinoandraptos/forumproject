package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// define DB as a pointer to a database object
var DB *sql.DB

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "220402"
	Dbname   = "forumdb"
)
