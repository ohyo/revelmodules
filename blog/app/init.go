package app

import (
	"fmt"
	"time"

	"github.com/revel/revel"
	"ohyo.network/modules/blog/app/models"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// date formatting
	revel.TemplateFuncs["formatDate"] = func(date time.Time) string {
		return date.Format("2006/01/02 03:04")
	}

	revel.TemplateFuncs["isAdmin"] = func(currentUser *models.User) bool {
		return currentUser != nil && currentUser.Role == "admin"
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	revel.OnAppStart(func() {
		fmt.Println("GOT PATH", revel.BasePath)
	})
}

// HeaderFilter is
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// TODO turn this into revel.HeaderFilter
	// should probably also have a filter for CSRF
	// not sure if it can go in the same filter or not

	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}