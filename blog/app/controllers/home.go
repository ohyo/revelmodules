package controllers

import "github.com/revel/revel"


// Home is
type Home struct {
	App
}

// Index is
func (ctrl Home) Index() revel.Result {
	log := ctrl.Log.New("BLOG", "HOME")
	log.Info("=========================")
	// ctrl.ViewArgs["ctrl"] = ctrl
	// ctrl.ViewArgs["title"] = "ooops"
	// ctrl.ViewArgs["moreStyles"] = []interface{}{}
	// ctrl.ViewArgs["moreScripts"] = []interface{}{}
	// ctrl.ViewArgs["RunMode"] = revel.RunMode
	
	return ctrl.RenderTemplate("Home/Index.jet.html")
}

// Example is
func (ctrl Home) Example() string {
	log := ctrl.Log.New("BLOG", "EXAMPLE")
	log.Info("EXAMPLE CALLED")
	return "CONTROLLER CALLED"
}
