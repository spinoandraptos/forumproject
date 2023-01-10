package models

import "time"

// a forum thread has an unique ID, title, content, category ID, author's name and author's ID
// the struct also reflects time of creation and updating of the thread
// note: while the thread struct contains the author's name, this information is not stored in threads table
// as it can simply be looked up through an inner join of tables

type Thread struct {
	ID             uint32 `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	AuthorID       uint32 `json:"authorid"`
	Authorusername string `json:"authorusername"`
	CategoryID     uint32 `json:"categoryid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
