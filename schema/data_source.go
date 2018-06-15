package schema

import (
	"github.com/savfx/savgo/util/convert"
)

type DataSource interface {
	IsForm() bool
	GetFormObject() *convert.FormObject
	GetFormArray() *convert.FormArray
	GetObjectAccess() *convert.ObjectAccess
	GetArrayAccess() * convert.ArrayAccess
}

type DefaultDataSource struct {
	FormObject *convert.FormObject
	FormArray *convert.FormArray
	ObjectAccess *convert.ObjectAccess
	ArrayAccess * convert.ArrayAccess
}

type Unmarshal func (data []byte, v interface{}) error

func NewJsonDataSource (data []byte, unmarshal Unmarshal) * DefaultDataSource {
	if len(data) > 0 {
		if data[0] == '{' {
			var src map[string]interface{}
			err := unmarshal(data, &src)
			if err == nil {
				return &DefaultDataSource{
					ObjectAccess: convert.NewObjectAccess(src),
				}
			}
		} else if data[0] == '[' {
			var src []interface{}
			err := unmarshal(data, &src)
			if  err == nil {
				return &DefaultDataSource{
					ArrayAccess: convert.NewArrayAccess(src),
				}
			}
		}
	}
	return nil
}

func NewFormDataSource (values map[string][]string) * DefaultDataSource {
	return &DefaultDataSource{
		FormObject: convert.NewFormObject(values),
	}
}

func (ctx DefaultDataSource) IsForm() bool {
	return ctx.FormArray != nil && ctx.FormObject != nil
}

func (ctx DefaultDataSource) GetFormObject() *convert.FormObject {
	return ctx.FormObject
}

func (ctx DefaultDataSource) GetFormArray() *convert.FormArray {
	return  ctx.FormArray
}

func (ctx DefaultDataSource) GetObjectAccess() *convert.ObjectAccess {
	return  ctx.ObjectAccess
}

func (ctx DefaultDataSource) GetArrayAccess() * convert.ArrayAccess {
	return  ctx.ArrayAccess
}
