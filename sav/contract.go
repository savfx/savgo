package sav

import "github.com/savfx/savgo/router"

type BaseContract struct {
	app Application
	name string
	modals map[string]Modal
	router * router.Router
	options map[string]interface{}
}

func (ctx BaseContract) GetOptions () map[string]interface{} {
	return  ctx.options
}

func (ctx * BaseContract) UpdateOptions(options map[string]interface{}) {
	if options != nil {
		for name, value := range options {
			ctx.options[name] = value
		}
		ctx.app.SyncContract(ctx)
	}
}

func (ctx BaseContract) GetName() string {
	return ctx.name
}

func (ctx BaseContract) GetModal(name string) Modal {
	return ctx.modals[name]
}

func (ctx BaseContract) GetRouter() * router.Router {
	return ctx.router
}

func (ctx BaseContract) SetJsonRoutes(contract string) {
	ctx.router.Load(contract)
}

func (ctx BaseContract) Fetch(modalName string, actionName string, handler DataHandler) (res Response, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	res, err = ctx.app.Fetch(ctx.GetAction(modalName, actionName), handler)
	if err == nil {
		handler.ParseOutput(res.GetData())
	}
	return res, err
}

func (ctx BaseContract) GetAction(modalName string, actionName string) Action {
	if modal, ok := ctx.modals[modalName]; ok {
		if action := modal.GetAction(actionName); action != nil {
			return action
		}
	}
	return nil
}

func (ctx * BaseContract) Init(app Application, name string) {
	ctx.name = name
	ctx.app = app
	ctx.modals = make(map[string]Modal)
	ctx.router = router.Create(&router.Option{
		Prefix    : "",
		Sensitive : false,
		Method    : "POST",
		CaseType  : "camel",
	})
}

func NewBaseContract(app Application, name string) *BaseContract {
	res := &BaseContract{}
	res.Init(app, name)
	return res
}

func (ctx * BaseContract) DefineModal (name string, actions map[string]ActionHandler) *BaseModal{
	res := NewBaseModal(ctx, name)
	if actions != nil {
		res.DefineActions(actions)
	}
	ctx.modals[name] = res
	return res
}
