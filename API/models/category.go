package models

// a forum category has an unique ID, title and description
type category struct {
	ID          uint32 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
