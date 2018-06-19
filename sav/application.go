package sav

import (
  "github.com/savfx/savgo/util/request"
  "github.com/savfx/savgo/util/convert"
  "time"
)

type BaseApplication struct {
  requests map[string]*request.Request
  contracts map[string]Contract
}

func (ctx BaseApplication) Fetch(action Action, handler DataHandler) (Response, error){
  name := action.GetContract().GetName()
  contract := ctx.contracts[name]
  actionName := action.GetModal().GetName() + action.GetName()
  router := contract.GetRouter()
  route := router.GetActionRoute(actionName)
  if len(route.Route.Keys) > 0 {

  }
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
