package models

import (
	"time"
)

// Comment is
type Comment struct {
	ID        int64
	Body      string
	Commenter string
	PostID    int // Foreign key for Post (belongs to)
	CreatedAt time.Time
	UpdatedAt time.Time
}
