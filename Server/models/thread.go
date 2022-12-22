package models

import (
	"log"
	"time"

	"github.com/spinoandraptos/forumproject/Server/database"
)

// a forum thread has an unique ID, title, content, category ID, author, author's ID and records the time of its creation or update
type Thread struct {
	ID         uint32 `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorID   uint32 `json:"authorid"`
	CategoryID uint32 `json:"categoryid"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// to retrieve all created threads from the database
// we select every row from the threads table and sort by most recent thread on top
// then we store every row in a thread and append the thread to a slice of threads
// finally we and close the database connection
func RetrieveAllThreads() (threadSlice []Thread, err error) {
	posts, err := database.DB.Query("SELECT * FROM threads ORDER BY CreatedAt DESC")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	for posts.Next() {
		var thread Thread
		err = posts.Scan(&thread.ID, &thread.Title, &thread.Content, &thread.AuthorID, &thread.CategoryID, &thread.CreatedAt, &thread.UpdatedAt)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		threadSlice = append(threadSlice, thread)
	}
	posts.Close()
	return
}
