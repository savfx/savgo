package sav

import (
	"github.com/savfx/savgo/util/convert"
	"github.com/savfx/savgo/router"
)

type DataSource interface {
	IsForm() bool
	GetFormObject() *convert.FormObject
	GetFormArray() *convert.FormArray
	GetObjectAccess() *convert.ObjectAccess
	GetArrayAccess() * convert.ArrayAccess
}

type DataHandler interface {
	GetInputValue () interface{}
	GetOutputValue () interface{}
	GetParams() map[string]interface{}
	ParseInput (ds DataSource)
	ParseOutput(ds DataSource)
}

type DataHandlerFactory func () DataHandler
type DataHandlerBinder func (handler DataHandler, value interface{})

type ActionHandler struct {
	Create DataHandlerFactory
	BindInput DataHandlerBinder
	BindOutput DataHandlerBinder
}

type Action interface {
	GetName() string
	GetModal() Modal
	GetContract() Contract
	GetHandler() *ActionHandler
}

type Modal interface {
	GetName() string
	GetAction(name string) Action
	GetContract() Contract
}

type Contract interface {
	GetName() string
	GetModal(name string) Modal
	GetRouter() *router.Router
	GetOptions () map[string]interface{}
	UpdateOptions(options map[string]interface{})
}

type Request interface {
	GetHeaders() map[string]string
	GetData() DataSource
}

type Response interface {
	GetStatus() int
	GetHeaders() map[string]string
	GetHeader(name string) string
	GetData() DataSource
}

type Application interface {
	SyncContract(contract Contract)
	Fetch(action Action, handler DataHandler) (Response, error)
}
