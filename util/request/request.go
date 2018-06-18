package request

import (
	"net/http"
	"net"
	"time"
	"io/ioutil"
	"strings"
	"context"
)

type Method uint8

const (
	GET  Method   = iota
	POST
	PUT
	DELETE
	PATCH
)

var MethodToString = map[Method]string{
	GET:     "GET",
	POST:    "POST",
	PUT:     "PUT",
	DELETE:  "DELETE",
	PATCH:   "PATCH",
}

type Headers map[string]string

type Request struct {
	client * http.Client
	options *Options
	connection *net.Conn
}

type Response struct {
	Headers Headers
	Body string
	Status string
	StatusCode int
}

type Options struct {
	Timeout time.Duration   // 超时时间, 单位是s
	KeepAlive time.Duration // 长链接持续时间, 单位是s
	BaseUrl string 			// 地址前缀
}

func NewRequest (options *Options) *Request {
	res := &Request{options: options}
	res.client = &http.Client{
		Transport: &http.Transport{
			DialContext: func (newContext context.Context, network, addr string) (net.Conn, error) {
				dial := net.Dialer{
					Timeout:   options.Timeout * time.Second,
					KeepAlive: options.Timeout * time.Second,
				}
				conn, err := dial.Dial(network, addr)
					if err != nil {
					return conn, err
				}
				res.connection = &conn
				return conn, err
			},
		},
	}
	return res
}

func toHttpHeaders (headers Headers, res * http.Header ) {
	for name, val := range headers {
		res.Add(name, val)
	}
}

func toLocalHeaders(headers http.Header) Headers {
	res := make(Headers)
	for name, val := range headers {
		res[name] = strings.Join(val, ",")
	}
	return  res
}

func (ctx Request) fetch (method Method, path string, headers Headers, data string) (*Response, error){
	if strings.Index(path, "//") == -1 {
		path = ctx.options.BaseUrl + path
	}
	var req *http.Request = nil
	var err error = nil
	switch method {
	case GET, DELETE:
		req, err = http.NewRequest(MethodToString[method], path,nil)
	case POST, PUT:
		req, err = http.NewRequest(MethodToString[method], path, strings.NewReader(data))
	}
	if err != nil {
		return nil, err
	}
	if headers != nil && len(headers) > 0 {
		toHttpHeaders(headers, &req.Header)
	}
	resp, err := ctx.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := &Response{
		StatusCode: resp.StatusCode,
		Status: resp.Status,
		Body: string(body),
		Headers: toLocalHeaders(resp.Header),
	}
	return res, nil
}

func (ctx Request) Get(path string, headers Headers) (*Response, error) {
	return ctx.fetch(GET, path, headers, "")
}

func (ctx Request) Post(path string, headers Headers, body string) (*Response, error) {
	return ctx.fetch(POST, path, headers, body)
}

func (ctx Request) PostForm(path string, headers Headers, body string) (*Response, error) {
	if headers == nil {
		headers = make(Headers)
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return ctx.fetch(POST, path, headers, body)
}

func (ctx Request) PostJson(path string, headers Headers, body string) (*Response, error) {
	if headers == nil {
		headers = make(Headers)
	}
	headers["Content-Type"] = "application/json"
	return ctx.fetch(POST, path, headers, body)
}
