package models

import "time"

// a thread comment has an unique ID, title, content, thread ID, author, author's ID
type Comment struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Author    User   `json:"author"`
	AuthorID  uint32 `json:"authorid"`
	ThreadID  uint32 `json:"threadid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
