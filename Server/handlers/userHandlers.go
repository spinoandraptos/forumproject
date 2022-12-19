package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/helper"
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

func ViewUser(w http.ResponseWriter, r *http.Request) {

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	helper.Catch(err)

	var human models.User
	response := database.DB.QueryRow("SELECT * FROM users WHERE ID = $1", userid)
	err = response.Scan(&human.ID, &human.Username, &human.Password, &human.CreatedAt, &human.UpdatedAt)
	helper.Catch(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(human)

	t, err := template.ParseFiles()
	helper.Catch(err)
	t.Execute(w, nil)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles()
	helper.Catch(err)
	t.Execute(w, nil)
}

func UserAuthentication(w http.ResponseWriter, r *http.Request) {

	var human models.User
	var humanTest models.User
	json.NewDecoder(r.Body).Decode(&human)
	response := database.DB.QueryRow("SELECT * FROM users WHERE Username = $1", human.Username)
	err := response.Scan(&humanTest.ID, &humanTest.Username, &humanTest.Password, &humanTest.CreatedAt, &humanTest.UpdatedAt)
	if err != nil {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Unable to Find User :(")
	}
	if human.Password == humanTest.Password {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Logged In Successfully!", "jwt": "12323432"})
	} else {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Login Failed, Please Check Password is Correct :(")
	}

	w.Write([]byte("HELLO!"))
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO!"))
}

/*
	-we first create a variable human of type user
	-then we decode the HTTP request body into the human variable
	-after which we execute the insertion of information stored in human to the database
	-createdat and updatedat timing of a new account should all be the current time of execution
	-finally we send a message througn ResponseWriter to tell the client that creation is complete
*/

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var human models.User
	json.NewDecoder(r.Body).Decode(&human)

	response, err := database.DB.Exec("INSERT INTO users (ID, Username, Password, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5)", &human.ID, &human.Username, &human.Password, time.Now(), time.Now())
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "User Creation Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "User Created Successfully!", "jwt": "1234456"})
	}
}

/*
	-we first create a variable human of type user
	-then we fetch the value of the URL parameter (user id) using chi.URLParam
	-then we decode the HTTP request body into the human variable
	-after which we execute the updation of information stored in human to the row specified by userid
	-updatedat timing will be the current time of execution
	-finally we send a message througn ResponseWriter to tell the client that updation is complete
*/

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var human models.User

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	helper.Catch(err)

	json.NewDecoder(r.Body).Decode(&human)

	response, err := database.DB.Exec("UPDATE users SET Username = $2, Password = $3, UpdatedAt = $4 WHERE ID = $1", userid, &human.Username, &human.Password, time.Now())
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "User Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "User Updated Successfully!"})
	}
}

/*
	-we first create a variable human of type user
	-then we fetch the value of the URL parameter (user id) using chi.URLParam
	-then we decode the HTTP request body into the human variable
	-after which we execute the deletion of information stored at the row specified by userid
	-finally we send a message througn ResponseWriter to tell the client that deletion is complete
*/

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var human models.User

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	helper.Catch(err)

	json.NewDecoder(r.Body).Decode(&human)

	response, err := database.DB.Exec("DELETE * FROM users WHERE ID = $1", userid)
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "User Deletion Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "User Deleted Successfully!"})
	}
}
