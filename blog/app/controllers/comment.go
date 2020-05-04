package controllers

import (
	"github.com/ohyo/revelmodules/blog/app/models"
	"github.com/ohyo/revelmodules/blog/app/routes"
	"github.com/revel/revel"
)

// Comment is
type Comment struct {
	App
}

// CheckUser is
func (ctrl Comment) CheckUser() revel.Result {
	if ctrl.MethodName != "Destroy" {
		return nil
	}

	if ctrl.CurrentUser == nil {
		ctrl.Flash.Error("Please log in first")
		return ctrl.Redirect(BlogUser.Login)
	}

	if ctrl.CurrentUser.Role != "admin" {
		ctrl.Response.Status = 401 // Unauthorized
		ctrl.Flash.Error("You are not admin")
		return ctrl.Redirect(BlogUser.Login)
	}
	return nil
}

// Create is
func (ctrl Comment) Create(postID int, body, commenter string) revel.Result {
	comment := models.Comment{PostID: postID, Body: body, Commenter: commenter}
	ctrl.Txn.Create(&comment)
	return ctrl.Redirect(routes.Post.Show(postID))
}

// Destroy is
func (ctrl Comment) Destroy(postID, ID int) revel.Result {
	ctrl.Txn.Where("id = ?", ID).Delete(&models.Comment{})
	return ctrl.Redirect(routes.Post.Show(postID))
}
