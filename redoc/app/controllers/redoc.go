package controllers

import (
	"github.com/revel/revel"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
	"os"
)


// Redoc is
type Redoc struct {
	*revel.Controller
}

func init() {
	revel.OnAppStart(refreshSwagger)
}

func exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func refreshSwagger() {	
	if revel.DevMode && exists(revel.AppPath + "/init.go") {
		err := gen.New().Build(&gen.Config{
			SearchDir:          revel.AppPath,
			MainAPIFile:        "init.go",
			PropNamingStrategy: swag.CamelCase,
			OutputDir:          revel.BasePath + "/docs",
			ParseVendor:        false,
		})

		if err != nil {
			revel.AppLog.Error(err.Error())
		}
	}
}

// Index is
func (ctrl Redoc) Index() revel.Result {
	ctrl.ViewArgs["ctrl"] = ctrl
	return ctrl.RenderTemplate("Redoc/Index.html")
}

// SwaggerJSON is
func (ctrl Redoc) SwaggerJSON() revel.Result {
	filename := revel.BasePath + "/docs/swagger.json"
	ctrl.Response.ContentType = "application/json"
	return ctrl.RenderFileName(filename, revel.NoDisposition)
}

// SwaggerYaml is
func (ctrl Redoc) SwaggerYaml() revel.Result {
	filename := revel.BasePath + "/docs/swagger.yaml"
	ctrl.Response.ContentType = "application/xml"
	return ctrl.RenderFileName(filename, revel.NoDisposition)
}
