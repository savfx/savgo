package sav

import (
	"testing"
	"github.com/a8m/expect"
)

func TestBaseContract(t *testing.T) {
	expect := expect.New(t)
	ctx := NewBaseContract(nil, "project")
	ctx.DefineModal("Account", map[string]ActionHandler{
		"Login": {},
		"Register": {},
	})
	expect(ctx != nil).To.Be.True()
}
