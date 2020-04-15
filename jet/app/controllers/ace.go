package controllers

import "github.com/revel/revel"

type JetController struct {
	*revel.Controller
}

// Called to render the ace template inner
func (c *JetController) RenderAceTemplate(base, inner string) revel.Result {
	c.ViewArgs["ace-inner"] = inner
	return c.RenderTemplate(base)
}
