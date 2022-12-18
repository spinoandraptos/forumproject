package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/models"
)

func ViewComment(w http.ResponseWriter, r *http.Request) {

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		catch(err)
	}
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	if err != nil {
		catch(err)
	}
	var reply models.Comment
	response := database.DB.QueryRow("SELECT * FROM comments WHERE ID = $1 AND ThreadID = $2 AND CategoryID = $3", commentid, threadid, categoryid)
	err = response.Scan(&reply.ID, &reply.Content, &reply.AuthorID, &reply.ThreadID, &reply.CreatedAt, &reply.UpdatedAt)
	catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)

}

func CreateComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment
	json.NewDecoder(r.Body).Decode(&reply)

	response, err := database.DB.Exec("INSERT INTO comments (ID, Content, AuthorID, ThreadID, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5, $6)", &reply.ID, &reply.Content, &reply.AuthorID, &reply.ThreadID, time.Now(), time.Now())
	catch(err)

	rowsAffected, err := response.RowsAffected()
	catch(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsAffected)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		catch(err)
	}
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	if err != nil {
		catch(err)
	}
	json.NewDecoder(r.Body).Decode(&reply)

	response, err := database.DB.Exec("UPDATE comments SET Content = $4, UpdatedAt = $5 WHERE ID = $1 AND ThreadID = $2 AND CategoryID = $3", commentid, threadid, categoryid, &reply.Content, time.Now())
	catch(err)

	rowsAffected, err := response.RowsAffected()
	catch(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsAffected)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		catch(err)
	}
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	if err != nil {
		catch(err)
	}
	json.NewDecoder(r.Body).Decode(&reply)

	response, err := database.DB.Exec("DELETE * FROM threads WHERE ID = $1 AND ThreadID = $2 AND CategoryID = $3", commentid, threadid, categoryid)
	catch(err)

	rowsAffected, err := response.RowsAffected()
	catch(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsAffected)
}
