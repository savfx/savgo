package checker

import (
	"reflect"
	"strconv"
	"strings"
)
/*
 * 验证器
 */

type Checker struct {
	Paths []string
	Msgs []string
}

type Checkable interface {
	Check (t* Checker) error
}

func (self *Checker) Ensure(v bool) *Checker {
	if v {
		return self
	}
	panic(ErrIsNil)
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
	panic(ErrIsNil)
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

func (self *Checker) Check(checkable Checkable) *Checker {
	err := checkable.Check(self)
	if (err != nil) {
		panic(err)
	}
	return self
}

func (self *Checker) Field(field string, msg string) *Checker {
	self.Paths = append(self.Paths, field)
	self.Msgs = append(self.Msgs, msg)
	return self
}

func (self *Checker) Message(msg string) *Checker {
	size := len(self.Paths)
	if size > 0 {
		self.Msgs[size-1] = msg
	}
	return self
}

func (self *Checker) Index(index int) *Checker {
	self.Paths = append(self.Paths, strconv.FormatInt(int64(index), 10))
	self.Msgs = append(self.Msgs, "")
	return self
}

func (self *Checker) Pop() *Checker {
	size := len(self.Paths) - 1
	if size >= 0 {
		self.Paths = self.Paths[: size]
		self.Msgs = self.Msgs[: size]
	}
	return self
}

func (self Checker) Path () string {
	return strings.Join(self.Paths, ".")
}
