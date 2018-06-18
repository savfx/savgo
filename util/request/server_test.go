package request

import (
	"testing"
	"github.com/a8m/expect"
)

func TestNewTextServer(t *testing.T) {
	expect := expect.New(t)
	svr := NewTextServer("hello world")
	go svr.Start()
	defer svr.Close()
	res, _ := svr.Client.Get("/", nil)
	expect(res.Body).To.Equal("hello world")
}

func TestNewJsonServer(t *testing.T) {
	 expect := expect.New(t)
	 svr := NewJsonServer(map[string]interface{}{"a": "b"})
	 go svr.Start()
	 defer svr.Close()
	 res, _ := svr.Client.Get("/", nil)
	 expect(res.Body).To.Equal(`{"a":"b"}`)
}

func TestNewJsonServer2(t *testing.T) {
	expect := expect.New(t)
	svr := NewJsonServer([]interface{}{"a", "b"})
	go svr.Start()
	defer svr.Close()
	res, _ := svr.Client.Get("/", nil)
	expect(res.Body).To.Equal(`["a","b"]`)
	expect(len(svr.Headers)).To.Equal(0)
	expect(svr.Body).To.Equal("")
}
