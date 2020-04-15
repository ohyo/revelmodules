package controllers

import (
	"github.com/ohyo/revelmodules/blog/app/models"
)

// App is
type App struct {
	GormController
	CurrentUser *models.User
}
