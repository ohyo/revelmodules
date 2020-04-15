package controllers

import (
	"ohyo.network/modules/blog/app/models"
)

// App is
type App struct {
	GormController
	CurrentUser *models.User
}

