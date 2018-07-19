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

type GinContext struct {
	c * gin.Context
	h sav.DataHandler
	done bool
}

func (ctx * GinContext) Accept (data interface{}) {
	ctx.done = true
	if data != nil {
		ctx.h.SetOutputValue(data)
		ctx.c.JSON(200, ctx.h.GetOutputValue())
	}
}

func (ctx * GinContext) Reject(code int, object interface{}) {
	ctx.done = true
	ctx.c.JSON(code, object)
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
		gc := &GinContext{
			c: g,
			h: dataHandler,
		}
		handler(g, controller, dataHandler, gc)
		if gc.done == false {
			gc.Accept(dataHandler.GetOutputValue())
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
