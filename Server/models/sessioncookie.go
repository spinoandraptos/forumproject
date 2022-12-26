package models

import (
	"time"
)

type Session struct {
	ID        int
	UUID      string
	Username  string
	CreatedAt time.Time
}
