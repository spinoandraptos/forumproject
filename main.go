//This file is the entry point for the app

package main

/*
	import standard library packages:
	the package which contains functions for formatting I/O text (fmt),
	the HTTP networking package(net/http),
	the logging package (log)
	the sql package with the corresponding postgres driver for working with database

	and third-party packages:
	the go-chi framework package
	tne go-chi middleware package
	the go-chi cors package
	other custom packages used in this project
*/

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/handlers"
	"github.com/spinoandraptos/forumproject/Server/helper"
)

// define the init function which connects to the postgresql database
// this will be run before main() as database package will be imported into main package
// then we define DB to point specifically to the forum database (specified by the info in psqlInfo which is sent to sql.Open)
// db.Ping will then attempt to open a connection with the database
// if error occurs, error message will be printed
func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", database.Host, database.Port, database.User, database.Password, database.Dbname)
	database.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	pingErr := database.DB.Ping()
	if err != nil {
		log.Fatal(pingErr)
	} else {
		fmt.Println("Successful Connection")
	}
}

/*
route definitions: all HTTP requests will be directed by the Chi Router to the respective handlers
if the URL path matches the format of /xxx/yyy, then the corresponding function will be called to send a response to the client
*/

func main() {
	//creation of a new router of name "router"
	router := chi.NewRouter()

	//use the go-chi inbuilt logger to log http requests and responses
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	//using the cors middleware to perform preflight CORS checks on the server side
	//this ensures that the server only permits browser requests fulfilling the below requirements which reduces potential misuse
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//Register endpoints (GET, POST, PUT, DELETE) with their respective paths
	//routes are grouped under 3 branches: one standalone for entry page, one for users and one for categories
	//route handler functions are defined under the respective handler files
	//it is to be noted that the handlers below are merely functions and do not implement the Handler interface
	//this is because using merely functions is clearer and simpler given we are not doing complex operations

	router.Route("/users", func(r chi.Router) {
		r.Post("/login", handlers.UserLogin)
		r.Post("/login/authenticate", handlers.UserAuthentication)
		r.Post("/{userid}/logout", handlers.UserLogout)
		r.Get("/{userid}", handlers.ViewUser)
		r.Post("/signup", handlers.CreateUser)
		r.Put("/{userid}", handlers.UpdateUser)
		r.Delete("/{userid}", handlers.DeleteUser)
	})

	router.Route("/", func(r chi.Router) {
		r.Get("/", handlers.ViewCategories)
		r.Get("/{categoryid}", handlers.ViewCategory)
		r.Get("/{categoryid}/threads", handlers.ViewThreads)
		r.Get("/{categoryid}/threads/{threadid}", handlers.ViewThread)
		r.Get("/{categoryid}/threads/{threadid}/comments", handlers.ViewComments)
		r.Get("/categories/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.ViewComment)

		router.Post("/{categoryid}/threads", handlers.CreateThread)
		router.Post("/{categoryid}/threads/{threadid}/comments", handlers.CreateComment)

		router.Put("/{categoryid}/threads/{threadid}", handlers.UpdateThread)
		router.Put("/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.UpdateComment)

		router.Delete("/{categoryid}/threads/{threadid}", handlers.DeleteThread)
		router.Delete("/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.DeleteComment)
	})

	//use router to start the server
	//if there is error starting server (error value is not nil), error message is printed
	err := http.ListenAndServe(":3000", router)
	helper.Catch(err)
}
