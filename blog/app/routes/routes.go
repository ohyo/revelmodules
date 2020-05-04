// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp



type tBlogUser struct {}
var BlogUser tBlogUser


func (_ tBlogUser) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BlogUser.CheckUser", args).URL
}

func (_ tBlogUser) Edit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BlogUser.Edit", args).URL
}

func (_ tBlogUser) Update(
		name string,
		oldPassword string,
		newPassword string,
		newPasswordConfirm string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "oldPassword", oldPassword)
	revel.Unbind(args, "newPassword", newPassword)
	revel.Unbind(args, "newPasswordConfirm", newPasswordConfirm)
	return revel.MainRouter.Reverse("BlogUser.Update", args).URL
}

func (_ tBlogUser) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BlogUser.Login", args).URL
}

func (_ tBlogUser) CreateSession(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("BlogUser.CreateSession", args).URL
}

func (_ tBlogUser) DestroySession(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BlogUser.DestroySession", args).URL
}


type tComment struct {}
var Comment tComment


func (_ tComment) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Comment.CheckUser", args).URL
}

func (_ tComment) Create(
		postID int,
		body string,
		commenter string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "postID", postID)
	revel.Unbind(args, "body", body)
	revel.Unbind(args, "commenter", commenter)
	return revel.MainRouter.Reverse("Comment.Create", args).URL
}

func (_ tComment) Destroy(
		postID int,
		ID int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "postID", postID)
	revel.Unbind(args, "ID", ID)
	return revel.MainRouter.Reverse("Comment.Destroy", args).URL
}


type tGormController struct {}
var GormController tGormController


func (_ tGormController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Begin", args).URL
}

func (_ tGormController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Rollback", args).URL
}

func (_ tGormController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Commit", args).URL
}


type tHome struct {}
var Home tHome


func (_ tHome) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Index", args).URL
}


type tPost struct {}
var Post tPost


func (_ tPost) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.CheckUser", args).URL
}

func (_ tPost) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.Index", args).URL
}

func (_ tPost) Show(
		ID int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ID", ID)
	return revel.MainRouter.Reverse("Post.Show", args).URL
}

func (_ tPost) Update(
		ID int,
		title string,
		body string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ID", ID)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "body", body)
	return revel.MainRouter.Reverse("Post.Update", args).URL
}

func (_ tPost) Create(
		title string,
		body string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "body", body)
	return revel.MainRouter.Reverse("Post.Create", args).URL
}

func (_ tPost) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.New", args).URL
}

func (_ tPost) Edit(
		ID int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ID", ID)
	return revel.MainRouter.Reverse("Post.Edit", args).URL
}

func (_ tPost) Destroy(
		ID int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ID", ID)
	return revel.MainRouter.Reverse("Post.Destroy", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


