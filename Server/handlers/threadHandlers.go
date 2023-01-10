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

// to view all threads under a category, we just retrieve category id from the URL sent in request
// then we look up all rows in threads table with a matching category id, along with the threads' author username, arranging them by time of creation in descending order
// then we scan the info for each thread into a thread struct within a slice of thread structs
// finally we send the entire slice back to client-side and close query connection

func ViewThreads(w http.ResponseWriter, r *http.Request) {

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	helper.Catch(err)
	var allthreads []models.Thread
	threads, err := database.DB.Query("SELECT threads.*, users.Username FROM threads INNER JOIN users ON threads.AuthorID=users.ID WHERE CategoryID = $1 ORDER BY CreatedAt DESC", categoryid)
	helper.Catch(err)
	for threads.Next() {
		var post models.Thread
		err = threads.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt, &post.Authorusername)
		helper.Catch(err)
		allthreads = append(allthreads, post)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allthreads)
	threads.Close()
}

// to view a specfic thread, retrieve both the thread id and corresponding category id from the URL sent in request
// and look up the row in threads table with a matching thread id and category id, along with the thread's author username
// then we scan all these info into a thread struct and send this struct back to client-side
// note: although just retrieving the thread id should suffice theoretically, I included category id for double-checking

func ViewThread(w http.ResponseWriter, r *http.Request) {

	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	helper.Catch(err)
	var post models.Thread
	thread := database.DB.QueryRow("SELECT threads.*, users.Username FROM threads INNER JOIN users ON threads.AuthorID=users.ID WHERE threads.ID = $1 AND threads.CategoryID = $2", threadid, categoryid)
	err = thread.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt, &post.Authorusername)
	helper.Catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

//SearchThread is essentially the same as ViewThread except it looks up the row in threads table by title instead of by id

func SearchThread(w http.ResponseWriter, r *http.Request) {

	var post models.Thread
	json.NewDecoder(r.Body).Decode(&post)
	thread := database.DB.QueryRow("SELECT * FROM threads WHERE threads.Title = $1", post.Title)
	err := thread.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt)
	helper.Catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// to create a thread, we retrieve the info to be inserted into threads table from the client-side request body
// then we insert the corresponding info into the table and check that we did modify the table

func CreateThread(w http.ResponseWriter, r *http.Request) {

	var post models.Thread
	json.NewDecoder(r.Body).Decode(&post)
	response, err := database.DB.Exec("INSERT INTO threads ( Title, Content, AuthorID, CategoryID, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5, $6)", &post.Title, &post.Content, &post.AuthorID, &post.CategoryID, time.Now(), time.Now())
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Thread Creation Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Thread Created Successfully!"})
	}
}

// to update a thread, we first retrieve both the thread id and corresponding category id from the URL sent in request
// we also retrieve the information that needs to be updated for the thread from request body
// then we update the row in threads table with a matching thread id and category id with the corresponding info
// finally we check that we have indeed modified the row

func UpdateThread(w http.ResponseWriter, r *http.Request) {

	var post models.Thread
	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	helper.Catch(err)
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
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

//this is a special case of UpdateThread where only the title needs to be changed

func UpdateThreadTitle(w http.ResponseWriter, r *http.Request) {

	var post models.Thread
	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	helper.Catch(err)
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	json.NewDecoder(r.Body).Decode(&post)
	response, err := database.DB.Exec("UPDATE threads SET Title = $3, UpdatedAt = $4 WHERE ID = $1 AND CategoryID = $2", threadid, categoryid, &post.Title, time.Now())
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Thread Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Thread Updated Successfully!"})
	}
}

//this is a special case of UpdateThread where only the content needs to be changed

func UpdateThreadContent(w http.ResponseWriter, r *http.Request) {

	var post models.Thread
	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	helper.Catch(err)
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	json.NewDecoder(r.Body).Decode(&post)
	response, err := database.DB.Exec("UPDATE threads SET Content = $3, UpdatedAt = $4 WHERE ID = $1 AND CategoryID = $2", threadid, categoryid, &post.Content, time.Now())
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Thread Update Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Thread Updated Successfully!"})
	}
}

// to delete a thread, we first retrieve both the thread id and corresponding category id from the URL sent in request
// then we delete the row in threads table with a matching thread id and category id and check that we have indeed removed the row

func DeleteThread(w http.ResponseWriter, r *http.Request) {

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	helper.Catch(err)
	threadid, err := strconv.Atoi(chi.URLParam(r, "threadid"))
	helper.Catch(err)
	response, err := database.DB.Exec("DELETE FROM threads WHERE ID = $1 AND CategoryID = $2", threadid, categoryid)
	helper.Catch(err)
	rowsAffected, err := response.RowsAffected()
	helper.Catch(err)
	if rowsAffected == 0 {
		helper.RespondwithERROR(w, http.StatusBadRequest, "Thread Deletion Failed :(")
	} else {
		helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Thread Deleted Successfully!"})
	}
}
