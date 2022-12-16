//This file is the entry point for the app

package main

/*
	import standard library packages:
	the package which contains functions for formatting I/O text (fmt),
	the HTTP networking package(net/http),
	the package that providesa platform-independent interface to operating system functionality (os),
	the package which contains functions for manipulating errors (errors)
	the logging package (log)
	the package which contains functions for manipulating UTF-8 encoded strings (strings)
	the package which provide functionality for measuring and displaying time (time)

	and one third-party package:
	the go-chi framework package
*/

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// define a function that catches errors by printing error message and panicks to stop execution
// the Recoverer middleware then recovers the server, logs the error with a stack trace, and sends a 500 Internal Server Error response to the client
func catch(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
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

	//Register endpoints (GET, POST, PUT, DELETE) with their respective paths
	//routes are grouped under 2 branches: one for users and one for categories
	//route handler functions are defined under the respective handler files

	router.Route("/users", func(r chi.Router) {
		r.Get("/login", userLogin)
		r.Get("/logout", userLogout)
		r.Get("/{userid}", viewUser)
		r.Post("/", createUser)
		r.Put("/{userid}", updateUser)
		r.Delete("/{userid}", deleteUser)
	})

	router.Route("/categories", func(r chi.Router) {
		r.Get("/", viewCategories)
		r.Get("/{categoryid}", viewCategory)
		r.Get("/{categoryid}/threads", viewThreads)
		r.Get("/{categoryid}/threads/{threadid}", viewThread)
		r.Get("/{categoryid}/threads/{threadid}/comments", viewComments)
		r.Get("/categories/{categoryid}/threads/{threadid}/comments/{commentid}", viewComment)

		router.Post("/{categoryid}/threads", createThread)
		router.Post("/{categoryid}/threads/{threadid}/comments", createComment)

		router.Put("/{categoryid}/threads/{threadid}", updateThread)
		router.Put("/{categoryid}/threads/{threadid}/comments/{commentid}", updateComment)

		router.Delete("/{categoryid}/threads/{threadid}", deleteThread)
		router.Delete("/{categoryid}/threads/{threadid}/comments/{commentid}", deleteComment)
	})

	//use router to start the server
	//if there is error starting server (error value is not nil), error message is printed and program exits
	err := http.ListenAndServe(":3000", router)
	catch(err)
}
