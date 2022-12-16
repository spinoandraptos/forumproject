package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/models"
)


/*
	-viewUser is called at the GET endpoint with URL parameter userid
	-we first fetch the value of the URL parameter (user id) using chi.URLParam (it returns the url parameter from a http.Request object )
	-strconv.Atoi is used to convert the returned string from URLParam into an int that is assigned to userid
	-then, based on the retrieved userid, we will find the row in table users that contains the matching id, whcich contains all info on the user we are looking for
	-$1 is a placeholder for userid in postgres notation
	-we scan the information of the retrieved row into a variable human of type user
	=then this information is sent back to the client through ResponseWriter
	-we first let the server inform the client that JSON data is being sent by setting content-type in header
	-then we encode the JSON information to be written to the client through ResponseWriter
	-we close the response and thus connection to the database at the very end so that it is ready for another connection later
*/

func viewUser(w http.ResponseWriter, r *http.Request) {

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	if err != nil {
		log.Println(err)
	}

	var human user
	response, err := DB.QueryRow("SELECT * FROM users WHERE ID = $1", userid).scan(&human.ID, &human.Username, &human.Password, &human.CreatedAt, &human.UpdatedAt)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(human)

	defer response.Close()
}

func userLogin(w http.ResponseWriter, r *http.Request) {

}

func userLogout(w http.ResponseWriter, r *http.Request) {

}

/*
	-we first create a variable human of type user
	-then we decode the HTTP request body into the human variable
	-after which we execute the insertion of information stored in human to the database
	-createdat and updatedat timing of a new account should all be the current time of execution
	-finally we send a message througn ResponseWriter to tell the client that creation is complete
*/

func createUser(w http.ResponseWriter, r *http.Request) {

	var human user
	json.NewDecoder(r.Body).Decode(&human)

	response, err := DB.Exec("INSERT INTO users (ID, Username, Password, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5)", &human.ID, &human.Username, &human.Password, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("User Created"))

	defer response.Close()
}

/*
	-we first create a variable human of type user
	-then we fetch the value of the URL parameter (user id) using chi.URLParam
	-then we decode the HTTP request body into the human variable
	-after which we execute the updation of information stored in human to the row specified by userid
	-updatedat timing will be the current time of execution
	-finally we send a message througn ResponseWriter to tell the client that updation is complete
*/

func updateUser(w http.ResponseWriter, r *http.Request) {

	var human user

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	if err != nil {
		log.Println(err)
	}

	json.NewDecoder(r.Body).Decode(&human)

	response, err := DB.Exec("UPDATE users SET Username = $2, Password = $3, UpdatedAt = $4 WHERE ID = $1", userid, &human.Username, &human.Password, time.Now())
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("User Updated"))
}

/*
	-we first create a variable human of type user
	-then we fetch the value of the URL parameter (user id) using chi.URLParam
	-then we decode the HTTP request body into the human variable
	-after which we execute the deletion of information stored at the row specified by userid
	-finally we send a message througn ResponseWriter to tell the client that deletion is complete
*/

func deleteUser(w http.ResponseWriter, r *http.Request) {

	var human user

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	if err != nil {
		log.Println(err)
	}

	json.NewDecoder(r.Body).Decode(&human)

	response, err := DB.Exec("DELETE * FROM users WHERE ID = $1", userid)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("User Deleted"))
}
