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
	Paths    []string
	Messages []string
}

type CheckHandler interface {
	Check(t *Checker) error
}

func (ctx *Checker) Ensure(v bool) *Checker {
	if v {
		return ctx
	}
	panic(ErrIsNil)
}

func (ctx *Checker) NotNull(v interface{}) *Checker {
	if v != nil {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Ptr:
			if !rv.IsNil() {
				return ctx
			}
		default:
			return ctx
		}
	}
	panic(ErrIsNil)
}

func (ctx *Checker) Exec(callback func()) (er error) {
	defer (func() {
		if err := recover(); err != nil {
			er = err.(error)
		}
	})()
	callback()
	return er
}

func (ctx *Checker) Check(handler CheckHandler) *Checker {
	err := handler.Check(ctx)
	if err != nil {
		panic(err)
	}
	return ctx
}

func (ctx *Checker) Field(field string, msg string) *Checker {
	ctx.Paths = append(ctx.Paths, field)
	ctx.Messages = append(ctx.Messages, msg)
	return ctx
}

func (ctx *Checker) Message(msg string) *Checker {
	size := len(ctx.Paths)
	if size > 0 {
		ctx.Messages[size-1] = msg
	}
	return ctx
}

func (ctx *Checker) Index(index int) *Checker {
	ctx.Paths = append(ctx.Paths, strconv.FormatInt(int64(index), 10))
	ctx.Messages = append(ctx.Messages, "")
	return ctx
}

func (ctx *Checker) Pop() *Checker {
	size := len(ctx.Paths) - 1
	if size >= 0 {
		ctx.Paths = ctx.Paths[:size]
		ctx.Messages = ctx.Messages[:size]
	}
	return ctx
}

func (ctx Checker) Path() string {
	return strings.Join(ctx.Paths, ".")
}
