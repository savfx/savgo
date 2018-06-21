package sav

import (
	"encoding/json"
	"github.com/savfx/savgo/router"
	"github.com/savfx/savgo/util/convert"
	"github.com/savfx/savgo/util/query"
	"github.com/savfx/savgo/util/request"
	"time"
)

type Middleware func(ctx *FetchContext) error

type BaseApplication struct {
	requests    map[string]*request.Request
	contracts   map[string]Contract
	middlewares []Middleware
}

func encodeQuery(value interface{}) (string, error) {
	text, err := json.Marshal(value)
	if err == nil {
		search := map[string]interface{}{}
		err = json.Unmarshal(text, &search)
		if err == nil {
			searchText := query.Encode(search)
			if searchText != "" {
				return searchText, nil
			}
		}
	}
	return "", err
}

type FetchContext struct {
	ContractName string
	ActionName   string
	Router       *router.Router
	Params       map[string]interface{}
	Headers      map[string]string
	Url          string
	Body         string
	Response     *request.Response
	Done         bool
}

func (ctx BaseApplication) Fetch(action Action, handler DataHandler) (Response, error) {
	name := action.GetContract().GetName()
	contract := ctx.contracts[name]
	actionName := action.GetModal().GetName() + action.GetName()
	route := contract.GetRouter().GetActionRoute(actionName)
	params := handler.GetParams()
	url := route.Route.Compile(params)
	body := ""
	var error error = nil
	if route.Method == router.GET || route.Method == router.DELETE {
		text, err := encodeQuery(handler.GetInputValue())
		if err == nil {
			if text != "" {
				url += "?" + text
			}
		} else {
			error = err
		}
	} else {
		if route.Form {
			text, err := encodeQuery(handler.GetInputValue())
			if err == nil {
				if text != "" {
					body = text
				}
			} else {
				error = err
			}
		} else {
			text, err := json.Marshal(handler.GetInputValue())
			if err == nil {
				body = string(text)
			} else {
				error = err
			}
		}
	}
	if error != nil {
		return nil, error
	}
	headers := map[string]string{}

	context := &FetchContext{
		ContractName: name,
		ActionName:   actionName,
		Params:       params,
		Headers:      headers,
		Url:          url,
		Body:         body,
	}

	for _, middleware := range ctx.middlewares {
		if err := middleware(context); err != nil {
			return nil, error
		}
	}
	var res *request.Response = nil
	if !context.Done {
		req := ctx.requests[name]
		switch route.Method {
		case router.GET:
			res, error = req.Get(context.Url, context.Headers)
		case router.POST:
			if route.Form {
				res, error = req.PostForm(context.Url, context.Headers, context.Body)
			} else {
				res, error = req.PostJson(context.Url, context.Headers, context.Body)
			}
		case router.PUT:
			if route.Form {
				res, error = req.PutForm(context.Url, context.Headers, context.Body)
			} else {
				res, error = req.PutJson(context.Url, context.Headers, context.Body)
			}
		case router.PATCH:
			if route.Form {
				res, error = req.PatchForm(context.Url, context.Headers, context.Body)
			} else {
				res, error = req.PatchJson(context.Url, context.Headers, context.Body)
			}
		case router.DELETE:
			res, error = req.Delete(context.Url, context.Headers)
		}
		context.Done = true
		context.Response = res
	}

	for n:= len(ctx.middlewares) -1 ; n >=0; n-- {
		if err := ctx.middlewares[n](context); err != nil {
			return nil, error
		}
	}

	if error != nil {
		return nil, error
	}

	response := &BaseResponse{
		StatusCode: res.StatusCode,
		Headers:    res.Headers,
		Body: res.Body,
		DataSource: NewJsonDataSource([]byte(res.Body), json.Unmarshal),
	}
	handler.ParseOutput(response.DataSource)

	return response, nil
}

func (ctx BaseApplication) SyncContract(contract Contract) {
	if _, ok := ctx.contracts[contract.GetName()]; !ok {
		ctx.contracts[contract.GetName()] = contract
		opts := convert.NewObjectAccess(contract.GetOptions())
		ctx.requests[contract.GetName()] = request.NewRequest(&request.Options{
			Timeout:   time.Duration(opts.GetInt64("Timeout")),
			KeepAlive: time.Duration(opts.GetInt64("KeepAlive")),
			BaseUrl:   opts.GetString("BaseUrl"),
		})
	} else {
		opts := convert.NewObjectAccess(contract.GetOptions())
		baseUrl := opts.GetString("BaseUrl")
		request := ctx.requests[contract.GetName()]
		request.SetBaseUrl(baseUrl)
	}
}

func (ctx *BaseApplication) AddMiddleware(middleware Middleware) {
	ctx.middlewares = append(ctx.middlewares, middleware)
}

func NewApplication() Application {
	res := &BaseApplication{
		contracts:   make(map[string]Contract),
		requests:    make(map[string]*request.Request),
		middlewares: make([]Middleware, 0),
	}
	return res
}
