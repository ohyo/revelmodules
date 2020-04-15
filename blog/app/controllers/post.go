package controllers

import (
	"github.com/revel/revel"
	"github.com/russross/blackfriday"
	"github.com/ohyo/revelmodules/blog/app/models"
	"github.com/ohyo/revelmodules/blog/app/routes"
	"html/template"
	"fmt"
)

// Post is
type Post struct {
	App
}

// CheckUser is
func (ctrl Post) CheckUser() revel.Result {

	ctrl.ViewArgs["ctrl"] = ctrl
	ctrl.ViewArgs["title"] = "ooops"
	ctrl.ViewArgs["Title"] = "ooops"
	ctrl.ViewArgs["moreStyles"] = []interface{}{}
	ctrl.ViewArgs["moreScripts"] = []interface{}{}
	ctrl.ViewArgs["RunMode"] = revel.RunMode

	switch ctrl.MethodName {
	case "Index", "Show":
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

// Index is
func (ctrl Post) Index() revel.Result {
	fmt.Println("HERE")
	var posts []models.Post
	ctrl.Txn.Order("created_at desc").Find(&posts)
	for i, p := range posts {
		posts[i].HTMLBody = template.HTML(string(blackfriday.MarkdownCommon([]byte(p.Body))))
	}
	fmt.Println("HERE", ctrl.ViewArgs)
	ctrl.ViewArgs["posts"] = posts
	return ctrl.RenderTemplate("Post/Index.jet.html")
}

// Show is
func (ctrl Post) Show(ID int) revel.Result {
	var post models.Post
	ctrl.Txn.First(&post, ID)
	ctrl.Txn.Where(&models.Comment{PostID: ID}).Find(&post.Comments)
	post.HTMLBody = template.HTML(string(blackfriday.MarkdownCommon([]byte(post.Body))))
	ctrl.ViewArgs["post"] = post
	return ctrl.RenderTemplate("Post/Show.jet.html")
}

// Update is
func (ctrl Post) Update(ID int, title, body string) revel.Result {
	var post models.Post
	ctrl.Txn.First(&post, ID)
	post.Title = title
	post.Body = body

	ctrl.Txn.Save(&post)
	return ctrl.Redirect(routes.Post.Show(ID))
}

// Create is
func (ctrl Post) Create(title, body string) revel.Result {
	post := models.Post{Title: title, Body: body}
	ctrl.Txn.Create(&post)
	return ctrl.Redirect(routes.Post.Show(int(post.ID)))
}

// New is
func (ctrl Post) New() revel.Result {
	post := models.Post{}
	ctrl.ViewArgs["post"] = post
	return ctrl.RenderTemplate("Post/Show.jet.html")
}

// Edit is
func (ctrl Post) Edit(ID int) revel.Result {
	var post models.Post
	ctrl.Txn.First(&post, ID)
	ctrl.ViewArgs["post"] = post
	return ctrl.RenderTemplate("Post/Edit.jet.html")
}

// Destroy is
func (ctrl Post) Destroy(ID int) revel.Result {
	ctrl.Txn.Where("post_id = ?", ID).Delete(&models.Comment{})
	ctrl.Txn.Where("id = ?", ID).Delete(&models.Post{})
	return ctrl.Redirect(routes.Post.Index())
}
