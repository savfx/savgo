package sav

type BaseModal struct {
	name string
	contract Contract
	actions map[string]Action
}

func (ctx BaseModal) GetName() string {
	return ctx.name
}

func (ctx BaseModal) GetContract() Contract {
	return ctx.contract
}

func (ctx BaseModal) GetAction(name string) Action {
	return ctx.actions[name]
}

func (ctx * BaseModal) DefineActions (actions map[string]ActionHandler) {
	for name, handler := range actions {
		item := NewBaseAction(ctx, name, handler)
		ctx.actions[name] = item
	}
}

func NewBaseModal(contract Contract, name string) *BaseModal {
	res := &BaseModal{
		contract:contract,
		name: name,
		actions: make(map[string]Action),
	}
	return res
}