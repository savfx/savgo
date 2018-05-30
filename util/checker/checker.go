package checker

import (
	"reflect"
)

/*
 * 验证器
 */

type Checker struct {
	field string
	msg   string
}

func (self *Checker) Field(field string, msg string) *Checker {
	self.field = field
	self.msg = msg
	return self
}

func (self *Checker) Ensure(v bool) *Checker {
	if v {
		return self
	}
	panic(ErrNotEnsure)
}

func (self *Checker) NotNull(v interface{}) *Checker {
	if v != nil {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Ptr:
			if !rv.IsNil() {
				return self
			}
		default:
			return self
		}
	}
	panic(ErrNotEnsure)
}

func (self *Checker) Exec(callback func()) (er error) {
	defer (func() {
		if err := recover(); err != nil {
			er = err.(error)
		}
	})()
	callback()
	return er
}
