package router

import (
	"testing"
	"github.com/a8m/expect"
)

func TestMatch(t *testing.T) {
	expect := expect.New(t)

	trustOpts := &ParseOption{
		Sensitive: true,
		End:       true,
		Strict:    true,
	}
	route := Parse("a", trustOpts)
	expect(route.Match("a")).To.Be.Map()
	expect(nil == route.Match("/a")).To.Be.True()
	expect(nil == route.Match("a/")).To.Be.True()

	route = Parse("a", &ParseOption{Sensitive: true, End: true, Strict: false})
	expect(route.Match("a")).To.Be.Map()
	expect(route.Match("a/")).To.Be.Map()
	expect(nil == route.Match("/a")).To.Be.True()

	route = Parse("/a", trustOpts)
	expect(route.Match("/a")).To.Be.Map()
	expect(route.Match("a") == nil).To.Be.True()
	expect(route.Match("a/") == nil).To.Be.True()

	route = Parse("a/", trustOpts)
	expect(route.Match("a/")).To.Be.Map()
	expect(route.Match("a") == nil).To.Be.True()
	expect(route.Match("/a") == nil).To.Be.True()

	route = Parse(":a", trustOpts)
	expect(route.Match("a")).To.Be.Map()
	expect(route.Match("/a") == nil).To.Be.True()
	expect(route.Match("a/") == nil).To.Be.True()

	route = Parse("/:a", trustOpts)
	expect(route.Match("/a")).To.Be.Map()
	expect(route.Match("a") == nil).To.Be.True()
	expect(route.Match("a/") == nil).To.Be.True()

	route = Parse(":a?", trustOpts)
	expect(route.Match("a")).To.Be.Map()
	expect(route.Match("")).To.Be.Map()
	expect(route.Match("/a") == nil).To.Be.True()
	expect(route.Match("a/") == nil).To.Be.True()

	route = Parse(":a/:b?", trustOpts)
	expect(route.Match("a")).To.Be.Map()
	expect(route.Match("a/b")).To.Be.Map()
	expect(route.Match("a/b/") == nil).To.Be.True()
	expect(route.Match("/a/b/") == nil).To.Be.True()

	route = Parse(":a/:b", trustOpts)
	expect(route.Match("a/b")).To.Be.Map()
	expect(route.Match("a/b/") == nil).To.Be.True()
	expect(route.Match("/a/b/") == nil).To.Be.True()

	route = Parse(":a-:b", trustOpts)
	expect(route.Match("a-b")).To.Be.Map()
	expect(route.Match("a-b/") == nil).To.Be.True()
	expect(route.Match("/a-b") == nil).To.Be.True()
	expect(route.Match("/a-b/") == nil).To.Be.True()

	route = Parse(":a-:b?", trustOpts)
	expect(route.Match("a-b")).To.Be.Map()
	expect(route.Match("a-")).To.Be.Map()
	expect(route.Match("a") == nil).To.Be.True()
	expect(route.Match("a-b/") == nil).To.Be.True()
	expect(route.Match("/a-") == nil).To.Be.True()
	expect(route.Match("/a-b") == nil).To.Be.True()
	expect(route.Match("/a-b/") == nil).To.Be.True()

	route = Parse(":a", &ParseOption{Sensitive: true, End: false, Strict: true})
	expect(route.Match("a")).To.Be.Map()
	expect(route.Match("a/")).To.Be.Map()
	expect(route.Match("a/b")).To.Be.Map()
	expect(route.Match("/a") == nil).To.Be.True()

	route = Parse("/home/:path?", &ParseOption{Sensitive: true, End: true, Strict: true})

	expect(route.Match("/home")).To.Be.Map()
	expect(route.Match("/home/a")).To.Be.Map()
	expect(route.Match("/HOME") == nil).To.Be.True()
	expect(route.Match("/HOME/a") == nil).To.Be.True()

	route = Parse("/home/:path?", &ParseOption{Sensitive: false, End: true, Strict: true})
	expect(route.Match("/home")).To.Be.Map()
	expect(route.Match("/home/a")).To.Be.Map()
	expect(route.Match("/HOME")).To.Be.Map()
	expect(route.Match("/HOME/a")).To.Be.Map()

}

func TestCompile(t *testing.T) {
	expect := expect.New(t)
	route := Parse("test", &ParseOption{})
	expect(route.Compile(map[string]interface{}{})).To.Equal("test")

	route = Parse("/test", &ParseOption{})
	expect(route.Compile(map[string]interface{}{})).To.Equal("/test")

	route = Parse(":a", &ParseOption{})
	expect(route.Compile(map[string]interface{}{"a": 1})).To.Equal("1")
	expect(route.Compile(map[string]interface{}{"a": "s"})).To.Equal("s")
	expect(route.CompileStrings(map[string]string{"a": "s b"})).To.Equal("s%20b")
	// 这个不支持
	// expect(route.Compile(map[string]interface{}{"a": "s:b"})).To.Equal("s%3Ab")
	expect(route.CompileStrings(map[string]string{"a": "s/b"})).To.Equal("s%2Fb")
	expect(route.Compile(map[string]interface{}{"a": "s;b"})).To.Equal("s%3Bb")
	expect(route.Compile(map[string]interface{}{"a": "s?b"})).To.Equal("s%3Fb")
	expect(route.Compile(map[string]interface{}{"a": true})).To.Equal("true")

	route = Parse("/:a", &ParseOption{})
	expect(route.Compile(map[string]interface{}{"a": 1})).To.Equal("/1")
	expect(route.Compile(map[string]interface{}{"a": "s"})).To.Equal("/s")
	expect(route.Compile(map[string]interface{}{"a": "s b"})).To.Equal("/s%20b")
	expect(route.Compile(map[string]interface{}{"a": true})).To.Equal("/true")

	route = Parse("/:a/:b", &ParseOption{})
	expect(route.Compile(map[string]interface{}{"a": 1, "b": 2})).To.Equal("/1/2")
	expect(route.Compile(map[string]interface{}{"a": "s", "b": "b"})).To.Equal("/s/b")
	expect(route.Compile(map[string]interface{}{"a": "s b", "b": "c"})).To.Equal("/s%20b/c")
	expect(route.Compile(map[string]interface{}{"a": true, "b": false})).To.Equal("/true/false")

	route = Parse("/:a?", &ParseOption{})
	expect(route.Compile(map[string]interface{}{})).To.Equal("")
	expect(route.Compile(map[string]interface{}{"a": 1})).To.Equal("/1")
	expect(route.Compile(map[string]interface{}{"a": "s"})).To.Equal("/s")
	expect(route.Compile(map[string]interface{}{"a": "s b"})).To.Equal("/s%20b")
	expect(route.Compile(map[string]interface{}{"a": true})).To.Equal("/true")

	route = Parse("/:a/:b?", &ParseOption{})
	expect(route.Compile(map[string]interface{}{"a": 1})).To.Equal("/1")
	expect(route.Compile(map[string]interface{}{"a": 1, "b": 2})).To.Equal("/1/2")
	expect(route.Compile(map[string]interface{}{"a": "s", "b": "b"})).To.Equal("/s/b")
	expect(route.Compile(map[string]interface{}{"a": "s b", "b": "c"})).To.Equal("/s%20b/c")
	expect(route.Compile(map[string]interface{}{"a": true, "b": false})).To.Equal("/true/false")

}
