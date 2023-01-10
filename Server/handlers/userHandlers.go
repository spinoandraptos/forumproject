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

// to view an user, retrieve the username from client-side request body
// and look up the row in users table with a matching username
// then we scan the row's info into an user struct and send this struct back to client-side

func ViewUser(w http.ResponseWriter, r *http.Request) {

	var human models.User
	json.NewDecoder(r.Body).Decode(&human)
	response := database.DB.QueryRow("SELECT * FROM users WHERE Username = $1", human.Username)
	err := response.Scan(&human.ID, &human.Username, &human.Password, &human.CreatedAt, &human.UpdatedAt)
	helper.Catch(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(human)
}

// this is a special case of ViewUser, where only the user id is sent back to client-side

func ViewUserID(w http.ResponseWriter, r *http.Request) {

	var human models.User
	json.NewDecoder(r.Body).Decode(&human)
	response := database.DB.QueryRow("SELECT * FROM users WHERE Username = $1", human.Username)
	err := response.Scan(&human.ID, &human.Username, &human.Password, &human.CreatedAt, &human.UpdatedAt)
	helper.Catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(human.ID)
}

// to create an user, we retrieve the info to be inserted into users table from the client-side request body
// then we insert the corresponding info into the table and check that we did modify the table

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

// to update an user, we first retrieve the user id from the URL sent in request
// we also retrieve the information that needs to be updated for the user from request body
// then we update the row in users table with a matching user id with the corresponding info
// finally we check that we have indeed modified the row

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

// this is a special case of UpdateUser where only the username is changed

func UpdateUsername(w http.ResponseWriter, r *http.Request) {

	var human models.User
	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	helper.Catch(err)
	json.NewDecoder(r.Body).Decode(&human)
	response, err := database.DB.Exec("UPDATE users SET Username = $2, UpdatedAt = $3 WHERE ID = $1", userid, &human.Username, time.Now())
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "User Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "User Updated Successfully!"})
	}
}

// this is a special case of UpdateUser where only the password is changed

func UpdateUserpassword(w http.ResponseWriter, r *http.Request) {

	var human models.User
	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	helper.Catch(err)
	json.NewDecoder(r.Body).Decode(&human)
	response, err := database.DB.Exec("UPDATE users SET Password = $2, UpdatedAt = $3 WHERE ID = $1", userid, &human.Password, time.Now())
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "User Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "User Updated Successfully!"})
	}
}

// to delete an user, we first retrieve the user id from the URL sent in request
// then we delete the row in users table with a matching user id and check that we have indeed removed the row

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	helper.Catch(err)
	response, err := database.DB.Exec("DELETE FROM users WHERE ID = $1", userid)
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "User Deletion Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "User Deleted Successfully!"})
	}
}
