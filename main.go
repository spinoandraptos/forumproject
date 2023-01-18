//This file is the entry point for the app

package main

/*
	import standard library packages:
	the package which contains functions for formatting I/O text (fmt),
	the HTTP networking package(net/http),
	the logging package (log)
	the package that allows for encoding and decoding of JSON
	the sql package with the corresponding postgres driver for working with database

	and third-party packages:
	the go-chi framework package
	tne go-chi middleware package
	the go-chi cors package
	the go-chi jwt package
	other custom packages used in this project
*/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/handlers"
	"github.com/spinoandraptos/forumproject/Server/helper"
	"github.com/spinoandraptos/forumproject/Server/models"
)

var authtoken *jwtauth.JWTAuth

const secretkey = "123abc"

// define the init function which connects to the postgresql database and assigns "authtoken" as a pointer to a JWT
// this will be run before main() as database package and "authtoken" will be imported into main package for use
// then we let DB point specifically to the forum database (specified by the info in psqlInfo which is sent to sql.Open)
// db.Ping will then attempt to open a connection with the database

func init() {
	var err error
	authtoken = jwtauth.New("HS256", []byte(secretkey), nil)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", database.Host, database.Port, database.User, database.Password, database.Dbname)
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

func main() {

	// creation of a new router of name "router"

	router := chi.NewRouter()

	// use the go-chi inbuilt logger to log http requests and responses
	// use the go-chi inbuilt recoverer to recover from panics

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	//using the cors middleware to perform preflight CORS checks on the server side
	//this ensures that the server only permits browser requests fulfilling the below requirements which reduces potential misuse

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://bopfishforum.onrender.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Set-Cookie", "Content-Type", "Authorisation", "Accept", "X-CSRF-Token", "Cookie"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//all HTTP requests will be directed by the Chi router to the respective handlers
	//if the URL path matches that specified, then the corresponding handler will be called to send a response to the client
	//routes are grouped under 2 branches: one for routes that can only be accessed by jwt-authenticated users
	//and one for public (unauthenticated) users

	router.Group(func(r chi.Router) {
		r.Post("/api/login", UserLogin)
		r.Post("/api/search", handlers.SearchThread)
		r.Post("/api/users/signup", handlers.CreateUser)
		r.Get("/api/categories", handlers.ViewCategories)
		r.Get("/api/{categoryid}", handlers.ViewCategory)
		r.Get("/api/{categoryid}/threads", handlers.ViewThreads)
		r.Get("/api/{categoryid}/threads/{threadid}", handlers.ViewThread)
		r.Get("/api/{categoryid}/threads/{threadid}/comments", handlers.ViewComments)
		r.Get("/api/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.ViewComment)
	})

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(authtoken))
		r.Use(jwtauth.Authenticator)
		r.Post("/api/users/{username}", handlers.ViewUser)
		r.Post("/api/users/{username}/id", handlers.ViewUserID)
		r.Post("/api/users/logout", UserLogout)
		r.Put("/api/users/{userid}", handlers.UpdateUser)
		r.Put("/api/users/{userid}/username", handlers.UpdateUsername)
		r.Put("/api/users/{userid}/password", handlers.UpdateUserpassword)
		r.Delete("/api/users/{userid}", handlers.DeleteUser)
		r.Post("/api/{categoryid}/threads", handlers.CreateThread)
		r.Post("/api/{categoryid}/threads/{threadid}/comments", handlers.CreateComment)
		r.Put("/api/{categoryid}/threads/{threadid}", handlers.UpdateThread)
		r.Put("/api/{categoryid}/threads/{threadid}/title", handlers.UpdateThreadTitle)
		r.Put("/api/{categoryid}/threads/{threadid}/content", handlers.UpdateThreadContent)
		r.Put("/api/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.UpdateComment)
		r.Delete("/api/{categoryid}/threads/{threadid}", handlers.DeleteThread)
		r.Delete("/api/{categoryid}/threads/{threadid}/comments/{commentid}", handlers.DeleteComment)
	})

	//use router to start the server on port 10000 as Render uses this port

	err := http.ListenAndServe(":10000", router)
	helper.Catch(err)
}

// to login, we first retrieve the login info from the client-side request body
// then we compare the info to that stored in the database
// if there is a match, we will encode the user info into a jwt that will be sent to the client-side as a cookie
// this cookie will help the user to keep logged in until they close the browser
// note: cookie must be named "jwt" to be recognised by jwtauth.Verifier

func UserLogin(w http.ResponseWriter, r *http.Request) {

	var human models.User
	var humanTest models.User
	json.NewDecoder(r.Body).Decode(&human)
	response := database.DB.QueryRow("SELECT * FROM users WHERE Username = $1", human.Username)
	err := response.Scan(&humanTest.ID, &humanTest.Username, &humanTest.Password, &humanTest.CreatedAt, &humanTest.UpdatedAt)
	helper.Catch(err)
	if human.Password == humanTest.Password {
		_, payloadclaims, _ := authtoken.Encode(map[string]interface{}{"username": &human.Username, "password": &human.Password})
		fmt.Printf("DEBUG: a sample jwt is %s\n\n", payloadclaims)
		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Name:     "jwt",
			Value:    payloadclaims,
		})
	} else {
		w.WriteHeader(404)
	}
}

// to log out, we simply terminate the session of the user by removing the cookie through setting MaxAge<0

func UserLogout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Name:     "jwt",
		Value:    "",
	})
}
