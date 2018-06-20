package router

import (
	"net/url"
	"regexp"
	"strconv"
)

type ParseOption struct {
	Sensitive bool
	End       bool
	Strict    bool
}

type Token struct {
	Name     string
	Text     string
	Prefix   string
	Optional bool
}

type Route struct {
	Regexp  string
	Tokens  []Token
	Keys    []string
	matcher *regexp.Regexp
}

func Parse(path string, opt *ParseOption) (route Route) {
	tokens := parseToken(path)
	regex := ""
	route.Tokens = tokens
	keys := make([]string, 0)
	for _, token := range tokens {
		if token.Name != "" {
			keys = append(keys, token.Name)
			if token.Prefix != "" {
				if token.Optional {
					regex = regex + `(?:/)?(?:([^/]+?))?`
				} else {
					regex = regex + `(?:/)(?:([^/]+?))`
				}
			} else {
				regex = regex + `(?:([^/]+?))`
				if token.Optional {
					regex = regex + `?`
				}
			}
		} else {
			regex = regex + escapeString([]byte(token.Text))
		}
	}
	if !opt.Strict {
		regex = regex + `/?`
	}
	if opt.End {
		regex = regex + `$`
	} else {
		regex = regex + `(?:/|$)?`
	}
	regex = "^" + regex
	if !opt.Sensitive {
		regex = `(?i)` + regex
	}
	route.Keys = keys
	route.Regexp = regex
	return
}

func (ctx *Route) Match(path string) map[string]string {
	if ctx.matcher == nil {
		ctx.matcher = regexp.MustCompile(ctx.Regexp)
	}
	args := ctx.matcher.FindAllStringSubmatch(path, -1)
	if args != nil {
		// dumpValue(args)
		list := args[0][1:len(args[0])]
		count := len(list)
		params := make(map[string]string, count)
		if count > 0 {
			for id, val := range list {
				if val != "" {
					params[ctx.Keys[id]] = val
				}
			}
		}
		return params
	}
	return nil
}

func (ctx *Route) Compile(params map[string]interface{}) string {
	data := make(map[string]string)
	for key, param := range params {
		str := ""
		switch v := param.(type) {
		case string:
			str = param.(string)
		case *string:
			str = *param.(*string)
		case int:
			str = strconv.FormatInt(int64(v), 10)
		case *int:
			str = strconv.FormatInt(int64(*v), 10)
		case bool:
			if param.(bool) {
				str = "true"
			} else {
				str = "false"
			}
		case *bool:
			if *param.(*bool) {
				str = "true"
			} else {
				str = "false"
			}
		}
		data[key] = str
	}
	return ctx.CompileStrings(data)
}

func (ctx *Route) CompileStrings(params map[string]string) (path string) {
	for _, token := range ctx.Tokens {
		if token.Name != "" {
			_, ok := params[token.Name]
			if !ok {
				if token.Optional {
					continue
				}
			}
			path = path + token.Prefix + url.PathEscape(params[token.Name])
		} else {
			path = path + token.Text
		}
	}
	return
}

var escapeRegexp = regexp.MustCompile(`([.+*?=^!:${}()[\]|/\\])`)

func escapeString(str []byte) string {
	return string(escapeRegexp.ReplaceAll(str, []byte("\\$1")))
}

var tokenRegexp = regexp.MustCompile(`((\/)?:(\w+)(\?)?)`)

func parseToken(path string) (tokens []Token) {
	parts := tokenRegexp.FindAllStringIndex(path, -1)
	pos := 0
	count := len(path)
	for i := 0; i < len(parts); i++ {
		mat := parts[i]
		offset := mat[0]
		end := mat[1]
		name := path[offset+1 : end]
		if pos < offset {
			tokens = append(tokens, Token{Text: path[pos:offset]})
		}
		prefix := ""
		if path[offset] != ':' {
			prefix = path[offset : offset+1]
			name = name[1:]
		}
		optional := false
		if path[end-1] == '?' {
			optional = true
			name = name[0 : len(name)-1]
		}
		tokens = append(tokens, Token{Name: name, Prefix: prefix, Optional: optional})
		pos = end
	}
	if pos < count {
		tokens = append(tokens, Token{Text: path[pos:count]})
	}
	return
}
