package sav

import (
	"github.com/a8m/expect"
	"testing"
)

func TestBaseContract(t *testing.T) {
	expect := expect.New(t)
	ctx := NewBaseContract(nil, "project", map[string]interface{}{})
	ctx.DefineModal("Account", map[string]ActionHandler{
		"Login":    {},
		"Register": {},
	})
	expect(ctx != nil).To.Be.True()
}
