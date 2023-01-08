package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// define DB as a pointer to a database object
// the DB is configured to be the one deployed on Render
var DB *sql.DB

const (
	Host     = "dpg-cerqoc6n6mphf4ufv3t0-a.singapore-postgres.render.com"
	Port     = 5432
	User     = "spino"
	Password = "wVQiWsmDnqOoRfFb5OvfnyDsg024undw"
	Dbname   = "forumdb_243o"
)

/*
const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "220402"
	Dbname   = "forumdb"
)
*/
