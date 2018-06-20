package sav

type BaseAction struct {
	name     string
	modal    Modal
	contract Contract
	handler  ActionHandler
}

func (ctx BaseAction) GetName() string {
	return ctx.name
}

func (ctx BaseAction) GetModal() Modal {
	return ctx.modal
}

func (ctx BaseAction) GetContract() Contract {
	return ctx.contract
}

func (ctx BaseAction) GetHandler() *ActionHandler {
	return &ctx.handler
}

func NewBaseAction(modal Modal, name string, handler ActionHandler) *BaseAction {
	res := &BaseAction{
		contract: modal.GetContract(),
		name:     name,
		modal:    modal,
		handler:  handler,
	}
	return res
}
