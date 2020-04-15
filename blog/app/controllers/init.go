package controllers

import (
	"github.com/revel/revel"
	"github.com/ohyo/revelmodules/blog/app/models"	
)

func (ctrl *App) setCurrentUser() revel.Result {
	log := ctrl.Log.New("Blog", "setCurrentUser")

	ctrl.ViewArgs["ctrl"] = ctrl
	ctrl.ViewArgs["title"] = "ooops"
	ctrl.ViewArgs["moreStyles"] = []interface{}{}
	ctrl.ViewArgs["moreScripts"] = []interface{}{}
	ctrl.ViewArgs["RunMode"] = revel.RunMode

	defer func() {
		log.Info("start " + ctrl.Action)

		if ctrl.CurrentUser != nil {
			ctrl.ViewArgs["currentUser"] = ctrl.CurrentUser
			log.Infof("current user: %q", ctrl.CurrentUser)
		} else {
			delete(ctrl.ViewArgs, "currentUser")
		}
	}()

	username, ok := ctrl.Session["username"]
	if !ok || username == "" {
		return nil
	}

	authKey, ok := ctrl.Session["authKey"]
	if !ok || authKey == "" {
		return nil
	}

	if match := revel.Verify(username.(string), authKey.(string)); match {
		var user models.User
		ctrl.Txn.Where(&models.User{Username: username.(string)}).First(&user)
		if &user != nil {
			ctrl.CurrentUser = &user
		}
	}

	return nil
}

func init() {
	revel.InterceptMethod((*App).setCurrentUser, revel.BEFORE)
	revel.InterceptMethod(BlogUser.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Post.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Comment.CheckUser, revel.BEFORE)
}
