package controllers

import "github.com/revel/revel"

type App struct {
	GormController
}

func (ctrl App) Index() revel.Result {
	return ctrl.Render()
}

func (ctrl App) ForwardAction(methodName string) {
	ctrl.MethodName = methodName
	ctrl.MethodType = ctrl.Type.Method(ctrl.MethodName)
	ctrl.Action = ctrl.Name + "." + ctrl.MethodName
}
