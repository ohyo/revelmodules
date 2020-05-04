package app

import (
	"io"
	"strings"

	module "github.com/ohyo/revelmodules/jet"

	"github.com/CloudyKit/jet/v3"
	"github.com/revel/revel"
)

// ENGINE_NAME is
const ENGINE_NAME = "jet"

// JetTemplate is Adapter for Go Templates.
type JetTemplate struct {
	*revel.TemplateView
	// engine *JetEngine
	template *jet.Template
	name     string
}

// Name is
func (tmpl JetTemplate) Name() string {
	revel.AppLog.Info("JET TEMPLATE NAME: " + tmpl.name)
	return tmpl.name
}

// // Content is
// func (tmpl JetTemplate) Content() []string {
// 	tmpl.
// 	return []string{}
// }

// Render is
// A bit trick of an implementation
// If the arg contains an ace_inner field then that will be used
// to fetch a new template
func (tmpl JetTemplate) Render(wr io.Writer, arg interface{}) error {
	// We can redirect this render to another template if the arguments contain ace_content in them

	// if argmap, ok := arg.(map[string]interface{}); ok {
	// 	if acecontentraw, ok := argmap["ace-inner"]; ok {
	// 		acecontent := acecontentraw.(string)
	// 		newtemplatename := tmpl.TemplateName + "-" + acecontent
	// 		// Now lookup the template again
	// 		if _, ok := tmpl.engine.templatesByName[newtemplatename]; !ok {
	// 			if inner, ok := tmpl.engine.templatesByName[acecontent]; !ok {
	// 				return fmt.Errorf("Inner content %s not found in ace templates", acecontent)
	// 			} else {
	// 				tmpl.engine.templatesByName[newtemplatename] = &JetTemplate{
	// 					// File:         tmpl.File,
	// 					// Inner:        inner.File,
	// 					engine:       tmpl.engine,
	// 					TemplateView: tmpl.TemplateView}
	// 			}

	// 		}
	// 		return tmpl.engine.templatesByName[newtemplatename].renderInternal(wr, arg)
	// 	}
	// }

	// revel.AppLog.Infof("JET ARG IS NULL %T", arg)
	// fmt.Println(reflect.TypeOf(tst))
	return tmpl.renderInternal(wr, arg)
}

func (tmpl JetTemplate) renderInternal(wr io.Writer, arg interface{}) error {
	// if tmpl.Template == nil {
	// 	// Compile the template first
	// 	if tmpl.Inner == nil {
	// 		tmpl.Inner = ace.NewFile("", nil)
	// 	}
	// 	source := ace.NewSource(tmpl.File, tmpl.Inner, tmpl.engine.files)
	// 	result, err := ace.ParseSource(source, tmpl.engine.Options)

	// 	if err != nil {
	// 		return err
	// 	}
	// 	if gtemplate, err := ace.CompileResult(tmpl.TemplateName, result, tmpl.engine.Options); err != nil {
	// 		return err
	// 	} else {
	// 		tmpl.Template = gtemplate
	// 	}
	// }
	return tmpl.template.Execute(wr, nil, arg)
}

// JetEngine is
type JetEngine struct {
	loader          *revel.TemplateLoader
	templates       map[string]*JetTemplate
	view            *jet.Set
	CaseInsensitive bool
}

// ConvertPath is
func (engine *JetEngine) ConvertPath(path string) string {
	if engine.CaseInsensitive {
		return strings.ToLower(path)
	}
	return path
}

// Handles is check that engine can handle it
func (engine *JetEngine) Handles(templateView *revel.TemplateView) bool {
	engine.view.AddPath(templateView.BasePath)
	can := revel.EngineHandles(engine, templateView)
	// if can {
	// 	revel.AppLog.Infof("JET CAN HANDLE: [%t] %s", can, templateView.TemplateName)
	// }
	return can
}

// ParseAndAdd is
func (engine *JetEngine) ParseAndAdd(templateView *revel.TemplateView) error {
	// revel.AppLog.Infof("JET PARSE AND ADD: [%s] [%s]", templateView.TemplateName, templateView.BasePath)
	// Jet templates must only render views specified for it (no trial and error)
	if templateView.EngineType != ENGINE_NAME {
		// revel.AppLog.Info("JET WRONG ENGINE TYPE: [" + templateView.FilePath + " " + templateView.TemplateName + "]")
		return &revel.Error{
			Title:       "Template Compilation Error",
			Path:        templateView.FilePath,
			Description: "Not correct template for engine",
			Line:        1,
			SourceLines: templateView.Content(),
		}
	}

	template, err := engine.view.GetTemplate(templateView.TemplateName)
	if err != nil {
		revel.AppLog.Error(err.Error())
		return err
	}
	// revel.AppLog.Info("JET ADD " + templateView.TemplateName)
	engine.templates[templateView.TemplateName] = &JetTemplate{
		// engine: engine,
		TemplateView: templateView,
		template:     template,
		name:         templateView.TemplateName,
	}

	return nil
}

// Lookup is
func (engine *JetEngine) Lookup(templateName string) revel.Template {
	if tpl, found := engine.templates[templateName]; found {
		// revel.AppLog.Info("JET FOUND: " + templateName)
		return tpl
	}
	return nil
}

// Name is
func (engine *JetEngine) Name() string {
	return ENGINE_NAME
}

// Event is
func (engine *JetEngine) Event(event revel.Event, i interface{}) {
	switch event {
	case revel.TEMPLATE_REFRESH_REQUESTED:
		if revel.RunMode == "dev" {
			engine.view.SetDevelopmentMode(true)
		}
	case revel.TEMPLATE_REFRESH_COMPLETED:
	}
}

func init() {
	// var root, _ = os.Getwd()
	err := revel.RegisterTemplateLoader(ENGINE_NAME, func(loader *revel.TemplateLoader) (revel.TemplateEngine, error) {
		revel.RevelLog.Debug("JET RegisterTemplateLoader")
		module.Engine = jet.NewHTMLSet()
		return &JetEngine{
			loader:    loader,
			view:      module.Engine,
			templates: map[string]*JetTemplate{},
		}, nil
	})

	if err != nil {
		revel.RevelLog.Error(err.Error())
	}
}
