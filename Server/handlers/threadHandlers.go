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

func ViewThreads(w http.ResponseWriter, r *http.Request) {

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		helper.Catch(err)
	}
	var allthreads []models.Thread
	threads, err := database.DB.Query("SELECT threads.*, users.Username FROM threads INNER JOIN users ON threads.AuthorID=users.ID WHERE CategoryID = $1 ORDER BY CreatedAt DESC", categoryid)
	if err != nil {
		helper.Catch(err)
	}
	for threads.Next() {
		var post models.Thread
		err = threads.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt, &post.Authorusername)
		if err != nil {
			helper.Catch(err)
		}
		allthreads = append(allthreads, post)
	}
	threads.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allthreads)
}

func ViewThread(w http.ResponseWriter, r *http.Request) {

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		helper.Catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		helper.Catch(err)
	}
	var post models.Thread
	response := database.DB.QueryRow("SELECT * FROM threads WHERE ID = $1 AND CategoryID = $2", threadid, categoryid)
	err = response.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CategoryID, &post.CreatedAt, &post.Authorusername)
	helper.Catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

func CreateThread(w http.ResponseWriter, r *http.Request) {

	var post models.Thread
	json.NewDecoder(r.Body).Decode(&post)

	response, err := database.DB.Exec("INSERT INTO threads (ID, Title, Content, AuthorID, CategoryID, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5, $6)", &post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CategoryID, time.Now(), time.Now())
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Thread Creation Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Thread Created Successfully!"})
	}
}

func UpdateThread(w http.ResponseWriter, r *http.Request) {

	var post models.Thread

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		helper.Catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		helper.Catch(err)
	}
	json.NewDecoder(r.Body).Decode(&post)

	response, err := database.DB.Exec("UPDATE threads SET Title = $3, Content = $4, UpdatedAt = $5 WHERE ID = $1 AND CategoryID = $2", threadid, categoryid, &post.Title, &post.Content, time.Now())
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Thread Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Thread Updated Successfully!"})
	}
}

func DeleteThread(w http.ResponseWriter, r *http.Request) {

	var post models.Thread

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	if err != nil {
		helper.Catch(err)
	}
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	if err != nil {
		helper.Catch(err)
	}
	json.NewDecoder(r.Body).Decode(&post)

	response, err := database.DB.Exec("DELETE * FROM threads WHERE ID = $1 AND CategoryID = $2", threadid, categoryid)
	helper.Catch(err)

	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)

	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Thread Deletion Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Thread Deleted Successfully!"})
	}
}
