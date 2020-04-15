package controllers

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ohyo/revelmodules/wiki/app/lib/akismet"
	"github.com/ohyo/revelmodules/wiki/app/lib/recaptcha"
	"github.com/ohyo/revelmodules/wiki/app/lib/wikihelper"
	"github.com/ohyo/revelmodules/wiki/app/models"

	"github.com/pmezard/go-difflib/difflib"
	"github.com/revel/revel"
)

type Page struct {
	App
	NeedRecaptcha bool
}

// Show page
func (ctrl Page) Show() revel.Result {
	pageName := ctrl.Params.Get("pageName")

	router := revel.MainRouter
	route := router.Route(ctrl.Request)
	controllertype := route.TypeOfController

	path := route.ModuleSource.Path
	importpath := route.ModuleSource.ImportPath
	namespace := route.ModuleSource.Namespace()

	fmt.Println("NAMESPACE", controllertype.Namespace, "PATH", path, "IMPORTPATH", importpath, "NAMESPACE", namespace, "NAME [", route.ModuleSource.Name, "]")

	// Home if no page name is specified
	if pageName == "" {
		pageName = "Home"
	}

	// Search by page name or page ID
	page := models.Page{}
	ctrl.db.Where("id = ?", pageName).Or("title = ?", pageName).First(&page)

	// If page ID is specified and exists, redirect to URL of page name
	id, _ := strconv.Atoi(pageName)
	if id != 0 && page.Id == id {
		return ctrl.Redirect("/page/" + wikihelper.UrlEncode(page.Title))
	}

	// Render Markdown
	body := page.Body
	html := wikihelper.Render(body)

	// Get revision number
	revision := 0
	ctrl.db.Model(models.Revision{}).Where("page_id = ?", page.Id).Count(&revision)

	// Get a list of recently registered pages
	recentCreatedPages := []models.Page{}
	ctrl.db.Order("created_at desc").Limit(10).Find(&recentCreatedPages)

	// Get recently updated page list
	recentUpdatedPages := []models.Page{}
	ctrl.db.Where("created_at != updated_at").Order("updated_at desc").Limit(10).Find(&recentUpdatedPages)

	for k, _ := range ctrl.ViewArgs {
		fmt.Println(k)
	}

	return ctrl.Render(pageName, body, html, page, revision, recentCreatedPages, recentUpdatedPages)
}

// Edit page
func (ctrl Page) Modify(pageName string) revel.Result {
	// Search by page name
	page := models.Page{}
	ctrl.db.Where("title = ?", pageName).First(&page)

	// Generate hash for collision detection
	hash := fmt.Sprintf("%x", sha1.Sum([]byte(page.Body)))

	// Take over Body
	if ctrl.Params.Get("page.Body") != "" {
		page.Body = ctrl.Params.Get("page.Body")
	}

	// CSRF
	nanoTime := time.Now().UnixNano()
	token := fmt.Sprintf("%x", sha1.Sum([]byte(strconv.FormatInt(nanoTime, 10))))
	ctrl.Session["token"] = token

	// Recaptcha
	recaptchaSiteKey := ""
	if recaptcha.IsEnabled() && (ctrl.NeedRecaptcha || recaptcha.IsAlways()) {
		recaptchaSiteKey, _ = recaptcha.GetKey()
	}

	return ctrl.Render(pageName, hash, page, token, recaptchaSiteKey)
}

// Register or update the page
func (ctrl Page) Save(pageName string) revel.Result {
	// Search by page name
	page := models.Page{}
	ctrl.db.Where("title = ?", pageName).First(&page)

	// Get the body sent by POST
	body := ctrl.Params.Get("page.Body")

	// Do not update if page exists but has no changes
	if page.Id > 0 && page.Body == body {
		return ctrl.Redirect("/page/" + wikihelper.UrlEncode(page.Title))
	}

	// CSRF
	token := ctrl.Params.Get("page.Token")
	if token != ctrl.Session["token"] {
		revel.AppLog.Info("DETECTED CSRF.")
		ctrl.ForwardAction("Modify")
		return ctrl.Modify(pageName)
	}

	// reCAPTCHA
	if recaptcha.IsAlways() && !recaptcha.Validate(ctrl.Controller) {
		revel.AppLog.Info("DETECTED INVALID CAPTCHA.")
		ctrl.ForwardAction("Modify")
		return ctrl.Modify(pageName)
	}

	// Akismet
	if !akismet.Validate(ctrl.Controller, body) {
		revel.AppLog.Info("DETECTED SPAM.")
		ctrl.NeedRecaptcha = true
		if recaptcha.IsEnabled() {
			if !recaptcha.IsAlways() && !recaptcha.Validate(ctrl.Controller) {
				revel.AppLog.Info("DETECTED INVALID CAPTCHA.")
				ctrl.ForwardAction("Modify")
				return ctrl.Modify(pageName)
			}
		} else {
			ctrl.ForwardAction("Modify")
			return ctrl.Modify(pageName)
		}
	}

	// Save the page
	page.Title = pageName
	page.Body = body
	ctrl.db.Save(&page)

	// Get latest revision
	previous := models.Revision{}
	ctrl.db.Where("page_id = ?", page.Id).Order("id desc").First(&previous)

	// Obtain the difference to count the added lines and deleted lines
	unifiedDiff := difflib.UnifiedDiff{
		A:       difflib.SplitLines(previous.Body),
		B:       difflib.SplitLines(page.Body),
		Context: 65535,
	}
	diffString, _ := difflib.GetUnifiedDiffString(unifiedDiff)
	diffLines := difflib.SplitLines(diffString)

	// Count added and deleted lines
	revision := models.Revision{}
	for i, line := range diffLines {
		if i > 2 {
			if strings.HasPrefix(line, "+") {
				revision.AddedLines++
			}
			if strings.HasPrefix(line, "-") {
				revision.DeletedLines++
			}
		}
	}

	// Save revision
	revision.Title = page.Title
	revision.Body = page.Body
	revision.PageId = page.Id
	ctrl.db.Save(&revision)

	return ctrl.Redirect("/page/" + wikihelper.UrlEncode(pageName))
}

// Display a list of page revisions
func (ctrl Page) Revisions(pageName string) revel.Result {
	// Search by page name
	page := models.Page{}
	ctrl.db.Where("title = ?", pageName).First(&page)

	revisions := []models.Revision{}
	ctrl.db.Where("page_id = ?", page.Id).Order("id desc").Find(&revisions)

	revisionSize := len(revisions)

	return ctrl.Render(pageName, revisions, revisionSize)
}

// Show diff between specified revision id and previous revision
// It is assumed to be requested by Ajax
func (ctrl Page) Diff(pageName string, revisionId string) revel.Result {
	// Search by page name
	page := models.Page{}
	ctrl.db.Where("title = ?", pageName).First(&page)

	// Get latest revision
	revision := models.Revision{}
	ctrl.db.Where("page_id = ? and id = ?", page.Id, revisionId).First(&revision)

	// Get revision immediately before latest revision
	previous := models.Revision{}
	ctrl.db.Where("page_id = ? and id < ?", page.Id, revisionId).Order("id desc").First(&previous)

	// Generate difference
	unifiedDiff := difflib.UnifiedDiff{
		A:       difflib.SplitLines(previous.Body),
		B:       difflib.SplitLines(revision.Body),
		Context: 65535,
	}
	diffString, _ := difflib.GetUnifiedDiffString(unifiedDiff)
	diffLines := difflib.SplitLines(diffString)

	// Remove unified diff header
	diffLines = diffLines[3:]

	// If the content before editing is empty, the first line is the difference that deletes the empty line, so delete
	if previous.Body == "" {
		diffLines = diffLines[1:]
	}
	diff := strings.Join(diffLines, "")
	diff = strings.TrimSpace(diff)

	// Render Markdown
	html := wikihelper.Render(revision.Body)

	return ctrl.Render(diff, revision, previous, pageName, html)
}
