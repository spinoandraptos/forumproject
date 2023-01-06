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
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	_ "github.com/lib/pq"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/handlers"
	"github.com/spinoandraptos/forumproject/Server/helper"
	"github.com/spinoandraptos/forumproject/Server/models"
)

// embed the react build files in main.go by specifying the directory where the build files are located
// then store the entire filesystem in variable named "reactfiles" which functions as a virtual filesystem

//go:embed frontend/build
var reactfiles embed.FS

var authtoken *jwtauth.JWTAuth

const secretkey = "123abc"

// define the init function which connects to the postgresql database and creates authtoken as a pointer to a JWT
// this will be run before main() as database package and authtoken will be imported into main package
// then we define DB to point specifically to the forum database (specified by the info in psqlInfo which is sent to sql.Open)
// db.Ping will then attempt to open a connection with the database
// if error occurs, error message will be printed
func init() {
	authtoken = jwtauth.New("HS256", []byte(secretkey), nil)
	var err error
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

/*
route definitions: all HTTP requests will be directed by the Chi Router to the respective handlers
if the URL path matches the format of /xxx/yyy, then the corresponding function will be called to send a response to the client
*/

func main() {

	//since all the files are located under frontend/build/static
	//we will serve the subdirectory of HOST/frontend/build/static directly using "filesystem" to avoid needing to access HOST/frontend/build/...
	//can access HOST/... directly
	filesystem, err := fs.Sub(reactfiles, "frontend/build")
	helper.Catch(err)
	newfilesystem := http.FS(filesystem)

	//creation of a new router of name "router"
	router := chi.NewRouter()

	//use the go-chi inbuilt logger to log http requests and responses
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

	//Register endpoints (GET, POST, PUT, DELETE) with their respective paths
	//routes are grouped under 2 branches: one for routes that can only be accessed by authenticated users
	//and one for public (unauthenticated) users
	//route handler functions are defined under the respective handler files
	//it is to be noted that the handlers below are merely functions and do not implement the Handler interface
	//this is because using merely functions is clearer and simpler given we are not doing complex operations
	router.Group(func(r chi.Router) {
		r.Handle("/", http.FileServer(newfilesystem))
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
		r.Post("/api/users/{username}", handlers.Viewuser)
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

	//use router to start the server
	//if there is error starting server (error value is not nil), error message is printed
	err = http.ListenAndServe(":3000", router)
	helper.Catch(err)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	var human models.User
	var humanTest models.User
	json.NewDecoder(r.Body).Decode(&human)
	response := database.DB.QueryRow("SELECT * FROM users WHERE Username = $1", human.Username)
	err := response.Scan(&humanTest.ID, &humanTest.Username, &humanTest.Password, &humanTest.CreatedAt, &humanTest.UpdatedAt)
	{
		helper.Catch(err)
	}
	if human.Password == humanTest.Password {
		_, payloadclaims, _ := authtoken.Encode(map[string]interface{}{"username": &human.Username, "password": &human.Password})
		fmt.Printf("DEBUG: a sample jwt is %s\n\n", payloadclaims)
		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Name:     "jwt", // Must be named "jwt" or else the token cannot be searched for by jwtauth.Verifier.
			Value:    payloadclaims,
		})
		decodedclaims, err := authtoken.Decode(payloadclaims)
		fmt.Printf("DEBUG: a sample jwt is %s\n\n", decodedclaims)
		helper.Catch(err)
	} else {
		w.WriteHeader(404)
	}
}

func UserLogout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1, // Delete the cookie.
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		Name:     "jwt",
		Value:    "",
	})
}
