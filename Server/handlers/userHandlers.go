package handlers

import (
	"encoding/json"
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

	var human models.User
	json.NewDecoder(r.Body).Decode(&human)
	response := database.DB.QueryRow("SELECT * FROM users WHERE Username = $1::varchar", human.Username)
	err := response.Scan(&human.ID, &human.Username, &human.Password, &human.CreatedAt, &human.UpdatedAt)
	helper.Catch(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(human.ID)
}

// the login function will take the juser info input of the request body and check if it matches the user data in the database
// if there is a match, we create a session for the user for session-based authentication
// after which we create a cookie with an unique UUID that will be stored in the browser and used for authentication
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
		jwttoken := models.Createtoken(&human.Username, &human.Password)
		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			Expires:  time.Now().Add(1 * time.Hour),
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Name:     "jwt", // Must be named "jwt" or else the token cannot be searched for by jwtauth.Verifier.
			Value:    jwttoken,
		})
	} else {
		w.WriteHeader(404)
	}
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	tokencookie, err := r.Cookie("jwt")
	if err != nil {
		helper.Catch(err)
	} else {
		cookievalue := tokencookie.Value
		var activesession models.Session
		response := database.DB.QueryRow("SELECT * FROM sessions WHERE UUID = $1", cookievalue)
		err := response.Scan(&activesession.ID, &activesession.UUID, &activesession.Username, &activesession.CreatedAt)
		if err != nil {
			helper.RespondwithERROR(w, http.StatusBadRequest, "Failed to Verify User")
		}
		deletion, err := database.DB.Exec("DELETE * FROM sessions WHERE UUID = $1", cookievalue)
		helper.Catch(err)

		rowsAffected, err := deletion.RowsAffected()
		helper.Catch(err)

		if rowsAffected == 0 {
			helper.RespondwithERROR(w, http.StatusBadRequest, "User Logout Failed :(")
		} else {
			helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "User Logged Out Successfully!"})
		}
	}
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

	response, err := database.DB.Exec("INSERT INTO users (Username, Password, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4)", &human.Username, &human.Password, time.Now(), time.Now())
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

/*
sessionuuid := uuid.NewString()
		fmt.Println(sessionuuid)
		response, err := database.DB.Exec("INSERT INTO sessions (Username, UUID, CreatedAt) VALUES ($1, $2, $3)", human.Username, sessionuuid, time.Now())
		if err != nil {
			w.WriteHeader(401)
		}
		rowsAffected, _ := response.RowsAffected()
		if rowsAffected != 0 {
			cookie := http.Cookie{
				Name:     "sessioncookie",
				Value:    sessionuuid,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			}
			http.SetCookie(w, &cookie)
			fmt.Println(cookie)
*/
