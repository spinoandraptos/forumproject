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
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/lib/pq"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/handlers"
	"github.com/spinoandraptos/forumproject/Server/helper"
)

var Authtoken *jwtauth.JWTAuth

const secretkey = "123abc"

// define the init function which connects to the postgresql database and creates authtoken as a pointer to a JWT
// this will be run before main() as database package and authtoken will be imported into main package
// then we define DB to point specifically to the forum database (specified by the info in psqlInfo which is sent to sql.Open)
// db.Ping will then attempt to open a connection with the database
// if error occurs, error message will be printed
func init() {
	Authtoken = jwtauth.New("HS256", []byte(secretkey), nil)
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
	//routes are grouped under 3 branches: one for routes that can only be accessed by authenticated users
	//one for branches under users and one for branches under categories
	//route handler functions are defined under the respective handler files
	//it is to be noted that the handlers below are merely functions and do not implement the Handler interface
	//this is because using merely functions is clearer and simpler given we are not doing complex operations
	router.Route("/", func(r chi.Router) {
		r.Get("/", handlers.ViewCategories)
		r.Get("/{categoryid}", handlers.ViewCategory)
		r.Get("/{categoryid}/threads", handlers.ViewThreads)
		r.Get("/{categoryid}/threads/{threadid}", handlers.ViewThread)
		r.Get("/{categoryid}/threads/{threadid}/comments", handlers.ViewComments)
		r.Get("/categories/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.ViewComment)
		r.Post("/{categoryid}/threads", handlers.CreateThread)
		r.Post("/{categoryid}/threads/{threadid}/comments", handlers.CreateComment)
		r.Put("/{categoryid}/threads/{threadid}", handlers.UpdateThread)
		r.Put("/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.UpdateComment)
		r.Delete("/{categoryid}/threads/{threadid}", handlers.DeleteThread)
		r.Delete("/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.DeleteComment)
	})

	router.Route("/users", func(r chi.Router) {
		r.Post("/login", handlers.UserLogin)
		r.Post("/signup", handlers.CreateUser)
		r.Get("/{userid}", handlers.ViewUser)
		r.Post("/logout", handlers.UserLogout)
		r.Put("/{userid}", handlers.UpdateUser)
		r.Delete("/{userid}", handlers.DeleteUser)
	})

	//use router to start the server
	//if there is error starting server (error value is not nil), error message is printed
	err := http.ListenAndServe(":3000", router)
	helper.Catch(err)
}
