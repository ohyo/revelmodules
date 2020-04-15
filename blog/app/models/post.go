package models

import (
	"html/template"
	"time"
)

// Post is
type Post struct {
	ID        int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment     // One-To-Many relationship (has many)
	HTMLBody  template.HTML `sql:"-"`
}
