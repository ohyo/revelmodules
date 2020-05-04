package controllers

import (
	"fmt"
	"os"

	"github.com/ohyo/revelmodules/redoc/app/swag"
	"github.com/ohyo/revelmodules/redoc/app/swag/gen"
	"github.com/revel/revel"
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
	fmt.Println("REFRESH SWAGGER")
	if revel.DevMode && exists(revel.AppPath+"/init.go") {
		err := gen.New().Build(&gen.Config{
			SearchDir: revel.AppPath,
			// SearchDirs:         []string{revel.AppPath, "/mnt/work/GO/src/github.com/ohyo/revelmodules/api/app"},
			MainAPIFile:        "swagger.go",
			PropNamingStrategy: swag.CamelCase,
			OutputDir:          revel.BasePath + "/public",
			ParseVendor:        false,
		})

		if err != nil {
			revel.AppLog.Error(err.Error())
		}
	} else {
		fmt.Println("SOME WRONG WITH SWAGGER")
	}
}

// Index is
func (ctrl Redoc) Index() revel.Result {
	ctrl.ViewArgs["ctrl"] = ctrl
	return ctrl.RenderTemplate("Redoc/Index.html")
}

// SwaggerJSON is
func (ctrl Redoc) SwaggerJSON() revel.Result {
	filename := revel.BasePath + "/public/swagger.json"
	ctrl.Response.ContentType = "application/json"
	return ctrl.RenderFileName(filename, revel.NoDisposition)
}

// SwaggerYaml is
func (ctrl Redoc) SwaggerYaml() revel.Result {
	filename := revel.BasePath + "/public/swagger.yaml"
	ctrl.Response.ContentType = "application/xml"
	return ctrl.RenderFileName(filename, revel.NoDisposition)
}
