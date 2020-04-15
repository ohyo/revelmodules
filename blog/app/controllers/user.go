package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"ohyo.network/modules/blog/app/models"
)

// User is
type BlogUser struct {
	App
}

// CheckUser is
func (ctrl BlogUser) CheckUser() revel.Result {

	ctrl.ViewArgs["ctrl"] = ctrl
	ctrl.ViewArgs["title"] = "ooops"
	ctrl.ViewArgs["moreStyles"] = []interface{}{}
	ctrl.ViewArgs["moreScripts"] = []interface{}{}
	ctrl.ViewArgs["RunMode"] = revel.RunMode

	switch ctrl.MethodName {
	case "Login", "CreateSession":
		return nil
	}

	if ctrl.CurrentUser == nil {
		ctrl.Flash.Error("Please log in first")
		return ctrl.Redirect(BlogUser.Login)
	}

	return nil
}

// Edit is
func (ctrl BlogUser) Edit() revel.Result {
	user := ctrl.CurrentUser
	ctrl.ViewArgs["user"] = user
	return ctrl.RenderTemplate("User/Edit.jet.html")
}

// Update is
func (ctrl BlogUser) Update(name, oldPassword, newPassword, newPasswordConfirm string) revel.Result {
	if err := bcrypt.CompareHashAndPassword(ctrl.CurrentUser.Password, []byte(oldPassword)); err != nil {
		ctrl.Flash.Error("Old password isn't valid.")
		return ctrl.Redirect(BlogUser.Edit)
	}

	var user models.User
	ctrl.Txn.First(&user, ctrl.CurrentUser.ID)
	user.Name = name

	if newPassword != "" && newPasswordConfirm != "" {
		if newPassword == newPasswordConfirm {
			bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
			user.Password = bcryptPassword
		} else {
			ctrl.Flash.Error("Password doesn't match the confirmation.")
			return ctrl.Redirect(BlogUser.Edit)
		}
	}

	ctrl.Txn.Save(&user)
	return ctrl.Redirect(Home.Index)
}

// Login is
func (ctrl BlogUser) Login() revel.Result {
	return ctrl.RenderTemplate("User/Login.jet.html")
}

// CreateSession is
func (ctrl BlogUser) CreateSession(username, password string) revel.Result {
	var user models.User
	ctrl.Txn.Where(&models.User{Username: username}).First(&user)

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err == nil {
		authKey := revel.Sign(user.Username)
		ctrl.Session["authKey"] = authKey
		ctrl.Session["username"] = user.Username
		ctrl.Session["userId"] = string(user.ID)
		if user.Role == "admin" {
			ctrl.Session["isAdmin"] = "true"
		}
		ctrl.Flash.Success("Welcome, " + user.Name)
		return ctrl.Redirect(Post.Index)
	}

	// clear session
	for k := range ctrl.Session {
		delete(ctrl.Session, k)
	}
	ctrl.Flash.Out["username"] = username
	ctrl.Flash.Error("Login failed")
	return ctrl.Redirect(Home.Index)
}

// DestroySession is
func (ctrl BlogUser) DestroySession() revel.Result {
	// clear session
	for k := range ctrl.Session {
		delete(ctrl.Session, k)
	}
	return ctrl.Redirect(Home.Index)
}
