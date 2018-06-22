package sgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/savfx/savgo/sav"
)

type Controller struct {
	Ctx *gin.Context
}

type IController interface {
	GetContext() interface{}
}

func (ctrl Controller) GetContext() interface{} {
	return ctrl.Ctx
}

func (ctrl Controller) SetContext(ctx * gin.Context) {
	ctrl.Ctx = ctx
}

type Binder interface {
	GetRawRequest() *http.Request
	RenderJson(code int, obj interface{})
}

type GinBinder struct {
	controller IController
}

func NewGinBinder(controller IController) Binder {
	return &GinBinder{controller: controller}
}

func (ctx GinBinder) GetRawRequest() *http.Request {
	return ctx.controller.GetContext().(*Controller).Ctx.Request
}

func (ctx GinBinder) RenderJson(code int, obj interface{}) {
	ctx.controller.GetContext().(*Controller).Ctx.JSON(code, obj)
}

type GinBindFunc func(c * gin.Context) Binder

type GinApplication struct {
	Engine * gin.Engine
	binder GinBindFunc
}

func (ctx GinApplication) Handle(g *gin.Context) {
	ctx.binder(g)
}

type RouteActionHandler func (ctrl IController, handler sav.DataHandler)

func (ctx GinApplication) GET(path string, handler RouteActionHandler) {
	//ctx.Engine.GET(path, ctx.MakeHandle())
}

func NewGinApplication (engine *gin.Engine, binder GinBindFunc) GinApplication{
	res := GinApplication{
		Engine: engine,
		binder: binder,
	}
	return res
}
