package sgin

import (
	"github.com/gin-gonic/gin"
	"github.com/savfx/savgo/sav"
	"github.com/savfx/savgo/router"
	"strings"
	"encoding/json"
)

type ControllerFactory func(c * gin.Context) sav.Controller

type GinApplication struct {
	engine * gin.Engine
	factory ControllerFactory
	contract sav.Contract
}

func (ctx GinApplication) MakeHandle(handler sav.RouteActionHandler, factory *sav.ActionHandler) func(g *gin.Context) {
	return func(g *gin.Context) {
		controller := ctx.factory(g)
		dataHandler := factory.Create()
		g.Request.ParseForm()
		formData := g.Request.PostForm
		if formData != nil && len(formData) > 0 {
			dataHandler.ParseInput(sav.NewFormDataSource(formData))
		} else {
			data, err := g.GetRawData()
			if err == nil {
				dataHandler.ParseInput(sav.NewJsonDataSource(data, json.Unmarshal))
			}
		}
		params := make(map[string]interface{})
		for _, param := range g.Params {
			params[param.Key] = param.Value
		}
		dataHandler.SetParams(params)
		handler(g, controller, dataHandler)
		data := dataHandler.GetOutputValue()
		if data != nil {
			g.JSON(200, data)
		}
	}
}

func (ctx GinApplication) Handle(modal, action string, handler sav.RouteActionHandler) {
	r := ctx.contract.GetRouter()
	route := r.GetActionRoute(modal + action)
	rawPath := r.GetPrefix() + route.Path
	path := strings.Replace(rawPath, "?", "", -1)
	ctx.engine.Handle(router.MethodToString[route.Method], path,
		ctx.MakeHandle(handler,
			ctx.contract.GetModal(modal).GetAction(action).GetHandler()))
}

func NewGinApplication (contract sav.Contract, engine *gin.Engine, factory ControllerFactory) GinApplication{
	res := GinApplication{
		engine: engine,
		factory: factory,
		contract: contract,
	}
	return res
}
