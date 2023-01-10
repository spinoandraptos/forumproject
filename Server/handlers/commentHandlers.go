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

// to view all comments under a thread, we just retrieve thread id from the URL sent in request
// then we look up all rows in comments table with a matching thread id, along with the comments' author username, arranging them by time of creation in descending order
// then we scan the info for each comment into a comment struct within a slice of comment structs
// finally we send the entire slice back to client-side and close query connection

func ViewComments(w http.ResponseWriter, r *http.Request) {
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	var allcomments []models.Comment
	comments, err := database.DB.Query("SELECT comments.*, users.Username FROM comments INNER JOIN users ON comments.AuthorID=users.ID WHERE ThreadID  = $1 ORDER BY CreatedAt DESC", threadid)
	helper.Catch(err)
	for comments.Next() {
		var reply models.Comment
		err = comments.Scan(&reply.ID, &reply.Content, &reply.AuthorID, &reply.ThreadID, &reply.CreatedAt, &reply.UpdatedAt, &reply.Authorusername)
		helper.Catch(err)
		allcomments = append(allcomments, reply)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allcomments)
	comments.Close()
}

// to view a specfic comment, retrieve both the comment id and corresponding thread id from the URL sent in request
// and look up the row in comments table with a matching comment id and thread id, along with the comment's author username
// then we scan all these info into a comment struct and send this struct back to client-side
// note: although just retrieving the comment id should suffice theoretically, I included thread id for double-checking

func ViewComment(w http.ResponseWriter, r *http.Request) {

	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	helper.Catch(err)
	var reply models.Comment
	response := database.DB.QueryRow("SELECT comments.*, users.Username FROM comments INNER JOIN users ON comments.AuthorID=users.ID WHERE comments.ID = $1 AND comments.ThreadID = $2", commentid, threadid)
	err = response.Scan(&reply.ID, &reply.Content, &reply.AuthorID, &reply.ThreadID, &reply.CreatedAt, &reply.UpdatedAt, &reply.Authorusername)
	helper.Catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)

}

// to create a comment, we retrieve the info to be inserted into comments table from the client-side request body
// then we insert the corresponding info into the table and check that we did modify the table

func CreateComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment
	json.NewDecoder(r.Body).Decode(&reply)
	response, err := database.DB.Exec("INSERT INTO comments (Content, AuthorID, ThreadID, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5)", &reply.Content, &reply.AuthorID, &reply.ThreadID, time.Now(), time.Now())
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Comment Creation Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Comment Created Successfully!"})
	}
}

// to update a comment, we first retrieve both the comment id and corresponding thread id from the URL sent in request
// we also retrieve the information that needs to be updated for the comment from request body
// then we update the row in comments table with a matching comment id and thread id with the corresponding info
// finally we check that we have indeed modified the row

func UpdateComment(w http.ResponseWriter, r *http.Request) {

	var reply models.Comment
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	helper.Catch(err)
	json.NewDecoder(r.Body).Decode(&reply)
	response, err := database.DB.Exec("UPDATE comments SET Content = $3, UpdatedAt = $4 WHERE ID = $1 AND ThreadID = $2", commentid, threadid, &reply.Content, time.Now())
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Comment Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Comment Updated Successfully!"})
	}
}

// to delete a comment, we first retrieve both the comment id and corresponding thread id from the URL sent in request
// then we delete the row in comments table with a matching comment id and thread id and check that we have indeed removed the row

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	commentid, err := strconv.Atoi(chi.URLParam(r, "commentid"))
	helper.Catch(err)
	response, err := database.DB.Exec("DELETE FROM comments WHERE ID = $1 AND ThreadID = $2", commentid, threadid)
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Comment Deletion Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Comment Deleted Successfully!"})
	}
}
