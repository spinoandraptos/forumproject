package models

import "time"

// a forum thread has an unique ID, title, content, category ID, author, author's ID and records the time of its creation or update
type Thread struct {
	ID         uint32 `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Author     User   `json:"author"`
	AuthorID   uint32 `json:"authorid"`
	CategoryID uint32 `json:"categoryid"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}
