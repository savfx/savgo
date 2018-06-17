package request

import (
	"testing"
	"github.com/a8m/expect"
)

func TestRequest_Get(t *testing.T) {
	expect := expect.New(t)
	svr := NewTextServer("")
	go svr.Start()
	defer svr.Close()
	svr.Client.Get("/", Headers{"a": "b"})
	expect(svr.Headers["A"]).To.Equal("b")
}

func TestRequest_Post(t *testing.T) {
	expect := expect.New(t)
	svr := NewTextServer("")
	go svr.Start()
	defer svr.Close()

	svr.Client.PostJson("/", nil, "{}")
	expect(svr.Body).To.Equal("{}")

	svr.Client.PostForm("/", nil, "a=b")
	expect(svr.Body).To.Equal("a=b")
}
