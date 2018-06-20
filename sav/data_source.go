package sav

import (
	"github.com/savfx/savgo/util/convert"
)

type BaseDataSource struct {
	FormObject *convert.FormObject
	FormArray *convert.FormArray
	ObjectAccess *convert.ObjectAccess
	ArrayAccess * convert.ArrayAccess
}

type Unmarshal func (data []byte, v interface{}) error

func NewJsonDataSource (data []byte, unmarshal Unmarshal) *BaseDataSource {
	if len(data) > 0 {
		if data[0] == '{' {
			var src map[string]interface{}
			err := unmarshal(data, &src)
			if err == nil {
				return &BaseDataSource{
					ObjectAccess: convert.NewObjectAccess(src),
				}
			}
		} else if data[0] == '[' {
			var src []interface{}
			err := unmarshal(data, &src)
			if  err == nil {
				return &BaseDataSource{
					ArrayAccess: convert.NewArrayAccess(src),
				}
			}
		}
	}
	return nil
}

func NewFormDataSource (values map[string][]string) *BaseDataSource {
	return &BaseDataSource{
		FormObject: convert.NewFormObject(values),
	}
}

func (ctx BaseDataSource) IsForm() bool {
	return ctx.FormArray != nil && ctx.FormObject != nil
}

func (ctx BaseDataSource) GetFormObject() *convert.FormObject {
	return ctx.FormObject
}

func (ctx BaseDataSource) GetFormArray() *convert.FormArray {
	return  ctx.FormArray
}

func (ctx BaseDataSource) GetObjectAccess() *convert.ObjectAccess {
	return  ctx.ObjectAccess
}

func (ctx BaseDataSource) GetArrayAccess() * convert.ArrayAccess {
	return  ctx.ArrayAccess
}

type BaseDataHandler struct {
	Params map[string]interface{}
}

func (ctx BaseDataHandler) GetParams () map[string]interface{} {
	return ctx.Params
}
