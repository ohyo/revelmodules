package models

import (
	"time"

	"github.com/revel/revel"
)

// Structure with page information
type Page struct {
	Id        int
	Title     string `sql:"size:255"`
	Body      string `sql:"size:16777215"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (page *Page) Validate(v *revel.Validation) {
	v.Check(
		page.Title,
		revel.Required{},
		revel.MaxSize{
			Max: 256,
		},
	)
	v.Check(
		page.Body,
		revel.Required{},
	)
}
