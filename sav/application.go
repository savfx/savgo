package sav

import (
  "github.com/savfx/savgo/util/request"
  "github.com/savfx/savgo/util/convert"
  "github.com/savfx/savgo/util/query"
  "github.com/savfx/savgo/router"
  "time"
  "encoding/json"
  "fmt"
)

type BaseApplication struct {
  requests map[string]*request.Request
  contracts map[string]Contract
}

func encodeQuery (value interface{}) (string, error) {
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

func (ctx BaseApplication) Fetch(action Action, handler DataHandler) (Response, error){
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
  fmt.Println(url, body, error)
  return nil, nil
}

func (ctx BaseApplication) SyncContract(contract Contract) {
  if  _, ok := ctx.contracts[contract.GetName()]; !ok {
    ctx.contracts[contract.GetName()] = contract
    opts := convert.NewObjectAccess(contract.GetOptions())
    ctx.requests[contract.GetName()] = request.NewRequest(&request.Options{
      Timeout: time.Duration(opts.GetInt64("Timeout")),
      KeepAlive: time.Duration(opts.GetInt64("KeepAlive")),
      BaseUrl: opts.GetString("BaseUrl"),
    })
  } else {
    opts := convert.NewObjectAccess(contract.GetOptions())
    baseUrl := opts.GetString("BaseUrl")
    request := ctx.requests[contract.GetName()]
    request.SetBaseUrl(baseUrl)
  }
}

func NewApplication() Application {
  res := &BaseApplication{
    contracts: make(map[string]Contract),
    requests: make(map[string]*request.Request),
  }
  return res
}
