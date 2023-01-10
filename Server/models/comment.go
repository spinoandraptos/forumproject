package models

import "time"

// a thread comment has an unique ID, content, thread ID, author's name and author's ID
// the struct also reflects time of creation and updating of comment
// note: while the comment struct contains the author's name, this information is not stored in comments table
// as it can simply be looked up through an inner join of tables

type Comment struct {
	ID             string `json:"id"`
	Content        string `json:"content"`
	AuthorID       uint32 `json:"authorid"`
	Authorusername string `json:"authorusername"`
	ThreadID       uint32 `json:"threadid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
