package models

import (
	"time"
)

// a forum user has an unique ID, username and password, and reflects time of creation and updating of user account
type User struct {
	ID        uint32 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
