package router

import (
	"github.com/savfx/savgo/util/strcase"
	"github.com/tidwall/gjson"
	"regexp"
	"strings"
)

type Method uint8

const (
	GET = iota
	POST
	PUT
	DELETE
	PATCH
	OPTIONS
	ANY
)

var MethodToString = map[Method]string{
	GET:     "GET",
	POST:    "POST",
	PUT:     "PUT",
	DELETE:  "DELETE",
	PATCH:   "PATCH",
	OPTIONS: "OPTIONS",
	ANY:     "ANY",
}

var StringToMethod = map[string]Method{
	"GET":     GET,
	"POST":    POST,
	"PUT":     PUT,
	"DELETE":  DELETE,
	"PATCH":   PATCH,
	"OPTIONS": OPTIONS,
	"ANY":     ANY,
}

type Option struct {
	Prefix    string
	Sensitive bool
	Method    string
	CaseType  string
	method    Method
	caseType  strcase.CaseType
}

type ModalRoute struct {
	Name      string
	Path      string
	Opts      gjson.Result
	ChildList map[Method][]*ActionRoute
	route     Route
}

type ActionRoute struct {
	Name   string
	Path   string
	Method Method
	Form   bool
	Opts   gjson.Result
	Modal  *ModalRoute
	Route  Route
}

type Router struct {
	opts           Option
	modalMap       map[string]*ModalRoute
	modalRoutes    []*ModalRoute
	actionRoutes   []*ActionRoute
	absoluteRoutes map[Method][]*ActionRoute
	actionMap      map[string]*ActionRoute
}

type MatchedRoute struct {
	Route  *ActionRoute
	Path   string
	Params map[string]string
}

func Create(opts *Option) *Router {
	router := &Router{}
	if opts.CaseType != "" {
		_, ok := strcase.StringToCaseType[opts.CaseType]
		if ok {
			opts.caseType = strcase.StringToCaseType[opts.CaseType]
		}
	}
	if opts.Method != "" {
		_, ok := StringToMethod[opts.Method]
		if ok {
			opts.method = StringToMethod[opts.Method]
		}
	}
	router.opts = *opts
	router.modalMap = make(map[string]*ModalRoute)
	router.absoluteRoutes = make(map[Method][]*ActionRoute)
	router.actionMap = make(map[string]*ActionRoute)
	return router
}

func (router *Router) Load(json string) {
	router.LoadGjson(gjson.Parse(json))
}

func (router *Router) LoadGjson(contract gjson.Result) {
	contract.Get("modals").ForEach(func(key gjson.Result, value gjson.Result) bool {
		router.createModalRoute(value, key.String())
		return true
	})
	contract.Get("actions").ForEach(func(key gjson.Result, value gjson.Result) bool {
		router.createActionRoute(value, key.String(), "")
		return true
	})
}

func (router *Router) MatchStringMethod(path string, method string) *MatchedRoute {
	method = strings.ToUpper(method)
	_, ok := StringToMethod[method]
	if ok {
		return router.Match(path, StringToMethod[method])
	}
	return nil
}

func (router *Router) Match(path string, method Method) *MatchedRoute {
	if method == OPTIONS {
		method = ANY
	}
	path = stripPrefix(path, router.opts.Prefix)
	matched := &MatchedRoute{
		Path: stripSuffix(path),
	}
	for _, actionRoute := range router.absoluteRoutes[method] {
		params := actionRoute.Route.Match(path)
		if params != nil {
			matched.Params = params
			matched.Route = actionRoute
			return matched
		}
	}
	for _, modalRoute := range router.modalRoutes {
		if modalRoute.route.Match(path) != nil {
			for _, actionRoute := range modalRoute.ChildList[method] {
				params := actionRoute.Route.Match(path)
				if params != nil {
					matched.Params = params
					matched.Route = actionRoute
					return matched
				}
			}
		}
	}
	return nil
}

func (router *Router) createModalRoute(opts gjson.Result, name string) {
	optName := opts.Get("name")
	if optName.Exists() {
		name = optName.String()
	}
	modalRoute := &ModalRoute{
		Name: strcase.Pascal(name),
		Path: strcase.Convert(name, router.opts.caseType),
		Opts: opts,
	}
	modalRoute.ChildList = make(map[Method][]*ActionRoute, 0)
	// 处理路径
	path := opts.Get("path")
	if path.Exists() {
		modalRoute.Path = path.String()
	}
	modalRoute.Path = normalPath("/" + modalRoute.Path)
	// 生成路由
	modalRoute.route = Parse(modalRoute.Path, &ParseOption{
		End:       false,
		Sensitive: router.opts.Sensitive,
	})
	// 添加到modalMap
	router.modalMap[name] = modalRoute
	id := opts.Get("id")
	if id.Exists() {
		router.modalMap[id.String()] = modalRoute
	}
	router.modalRoutes = append(router.modalRoutes, modalRoute)

	opts.Get("routes").ForEach(func(key gjson.Result, value gjson.Result) bool {
		router.createActionRoute(value, key.String(), name)
		return true
	})
}

func (router *Router) createActionRoute(opts gjson.Result, name string, modalName string) {
	optName := opts.Get("name")
	if optName.Exists() {
		name = optName.String()
	}
	optModalName := opts.Get("modal")
	if optModalName.Exists() {
		modalName = optModalName.String()
	}
	modal := router.modalMap[modalName]
	var method Method = GET
	optMethod := opts.Get("method")
	view := opts.Get("view")
	if optMethod.Exists() {
		method = StringToMethod[optMethod.String()]
	} else if view.Exists() {
		if view.Bool() {
			method = router.opts.method
		}
	} else {
		method = router.opts.method
	}
	actionRoute := &ActionRoute{
		Name:   modal.Name + strcase.Pascal(name),
		Opts:   opts,
		Method: method,
	}
	form := opts.Get("form")
	if form.Exists() {
		actionRoute.Form = form.Bool()
	}
	isAbsolute := false
	// 处理路径
	optPath := opts.Get("path")
	path := strcase.Convert(name, router.opts.caseType)
	if optPath.Exists() {
		pathVal := optPath.String()
		if len(pathVal) > 0 {
			if pathVal[0] == '/' {
				isAbsolute = true
				path = pathVal
			}
		}
		if !isAbsolute {
			path = modal.Path + "/" + pathVal
		}
	} else {
		path = modal.Path + "/" + path
	}
	path = normalPath(path)
	if path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}
	actionRoute.Path = path
	actionRoute.Modal = modal
	// 生成路由
	actionRoute.Route = Parse(path, &ParseOption{
		End:       true,
		Strict:    false,
		Sensitive: router.opts.Sensitive,
	})
	// 添加到容器
	if isAbsolute {
		router.absoluteRoutes[method] = append(router.absoluteRoutes[method], actionRoute)
		router.absoluteRoutes[ANY] = append(router.absoluteRoutes[ANY], actionRoute)
	} else {
		modal.ChildList[method] = append(modal.ChildList[method], actionRoute)
		modal.ChildList[ANY] = append(modal.ChildList[ANY], actionRoute)
	}
	router.actionRoutes = append(router.actionRoutes, actionRoute)
	router.actionMap[actionRoute.Name] = actionRoute
}

var normalRegexp = regexp.MustCompile(`/\/+/`)

func normalPath(str string) string {
	return string(normalRegexp.ReplaceAllString(str, "/"))
}

func stripPrefix(src string, prefix string) string {
	if len(prefix) > 0 && len(src) > 0 {
		pos := strings.Index(src, prefix)
		if pos == 0 || (pos == 1 && src[0] == '/') {
			start := pos + len(prefix)
			if start < len(src) {
				src = src[start:]
				if src[0] != '/' {
					return "/" + src
				}
			} else {
				return "/"
			}
		}
	}
	return src
}

func stripSuffix(path string) string {
	if len(path) > 0 {
		if path[len(path)-1] == '/' {
			return path[0 : len(path)-1]
		}
	}
	return path
}

func (router *Router) GetActionRoute(actionName string) *ActionRoute {
	if action, ok := router.actionMap[actionName]; ok {
		return action
	}
	return nil
}

func (router Router) GetPrefix() string {
	return router.opts.Prefix
}

func (router *Router) SetPrefix(prefix string) {
	router.opts.Prefix = prefix
}
