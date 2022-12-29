package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/helper"
	"github.com/spinoandraptos/forumproject/Server/models"
)

func ViewComments(w http.ResponseWriter, r *http.Request) {
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		helper.Catch(err)
	}
	var allcomments []models.Comment
	comments, err := database.DB.Query("SELECT comments.*, users.Username FROM comments INNER JOIN users ON comments.AuthorID=users.ID WHERE ThreadID  = $1 ORDER BY CreatedAt DESC", threadid)
	if err != nil {
		helper.Catch(err)
	}
	for comments.Next() {
		var reply models.Comment
		err = comments.Scan(&reply.ID, &reply.Content, &reply.AuthorID, &reply.ThreadID, &reply.CreatedAt, &reply.UpdatedAt, &reply.Authorusername)
		if err != nil {
			helper.Catch(err)
		}
		allcomments = append(allcomments, reply)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allcomments)
}

func ViewComment(w http.ResponseWriter, r *http.Request) {

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		helper.Catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		helper.Catch(err)
	}
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	if err != nil {
		helper.Catch(err)
	}
	var reply models.Comment
	response := database.DB.QueryRow("SELECT * FROM comments WHERE ID = $1 AND ThreadID = $2 AND CategoryID = $3", commentid, threadid, categoryid)
	err = response.Scan(&reply.ID, &reply.Content, &reply.AuthorID, &reply.ThreadID, &reply.CreatedAt, &reply.UpdatedAt)
	helper.Catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)

}

func CreateComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment
	json.NewDecoder(r.Body).Decode(&reply)

	response, err := database.DB.Exec("INSERT INTO comments (ID, Content, AuthorID, ThreadID, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5, $6)", &reply.ID, &reply.Content, &reply.AuthorID, &reply.ThreadID, time.Now(), time.Now())
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Comment Creation Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Comment Created Successfully!"})
	}
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		helper.Catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		helper.Catch(err)
	}
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	if err != nil {
		helper.Catch(err)
	}
	json.NewDecoder(r.Body).Decode(&reply)

	response, err := database.DB.Exec("UPDATE comments SET Content = $4, UpdatedAt = $5 WHERE ID = $1 AND ThreadID = $2 AND CategoryID = $3", commentid, threadid, categoryid, &reply.Content, time.Now())
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Comment Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Comment Updated Successfully!"})
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		helper.Catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		helper.Catch(err)
	}
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	if err != nil {
		helper.Catch(err)
	}
	json.NewDecoder(r.Body).Decode(&reply)

	response, err := database.DB.Exec("DELETE * FROM threads WHERE ID = $1 AND ThreadID = $2 AND CategoryID = $3", commentid, threadid, categoryid)
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Comment Deletion Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Comment Deleted Successfully!"})
	}
}
