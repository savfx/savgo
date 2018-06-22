package sgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/savfx/savgo/sav"
)

type GinController struct {
	Ctx *gin.Context
}

func (ctrl GinController) GetContext() interface{} {
	return ctrl.Ctx
}

func (ctrl GinController) SetContext(ctx * gin.Context) {
	ctrl.Ctx = ctx
}

type GinContextHandler struct {
	controller sav.Controller
}

func NewGinContextHandler(controller sav.Controller) sav.ContextHandler {
	return &GinContextHandler{controller: controller}
}

func (ctx GinContextHandler) GetRawRequest() *http.Request {
	return ctx.controller.GetContext().(*GinController).Ctx.Request
}

func (ctx GinContextHandler) RenderJson(code int, obj interface{}) {
	ctx.controller.GetContext().(*GinController).Ctx.JSON(code, obj)
}

type GinBindFunc func(c * gin.Context) sav.Controller

type GinApplication struct {
	Engine * gin.Engine
	binder GinBindFunc
	contract sav.Contract
}

func (ctx GinApplication) MakeHandle(handler RouteActionHandler, factory *sav.ActionHandler) func(g *gin.Context) {
	return func(g *gin.Context) {
		controller := ctx.binder(g)
		dataHandler := factory.Create()
		handler(controller, dataHandler)
	}
}

type RouteActionHandler func (ctrl sav.Controller, handler sav.DataHandler)

func (ctx GinApplication) GET(path string, modal, action string, handler RouteActionHandler) {
	ctx.Engine.GET(path, ctx.MakeHandle(handler, ctx.contract.GetModal(modal).GetAction(action).GetHandler()))
}

func NewGinApplication (engine *gin.Engine, binder GinBindFunc) GinApplication{
	res := GinApplication{
		Engine: engine,
		binder: binder,
	}
	return res
}
