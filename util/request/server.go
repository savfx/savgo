package request

import (
	"net/http"
	"net"
	"github.com/gin-gonic/gin/json"
	"fmt"
	"io/ioutil"
)

type ServerHandler struct {
	handler http.HandlerFunc
	listener net.Listener
	BaseUrl string
	Client *Request
	Body string
	Headers Headers
}

func NewHttpServer(handler http.HandlerFunc) *ServerHandler {
	listener, _ := net.Listen("tcp", "")
	port := listener.Addr().(*net.TCPAddr).Port
	url := fmt.Sprintf("http://localhost:%d", port)
	svr := &ServerHandler{
		handler:handler,
		listener:listener,
		BaseUrl: url,
		Client: NewRequest(&Options{
			BaseUrl: url,
		}),
	}
	return svr
}

func NewTextServer (body string) * ServerHandler {
	var res * ServerHandler = nil
	res = NewHttpServer(func (writer http.ResponseWriter, request *http.Request){
		res.processRequest(request)
		writer.Write([]byte(body))
	})
	return res
}

func NewJsonServer(data interface{}) * ServerHandler {
	var res *ServerHandler = nil
	res = NewHttpServer(func (writer http.ResponseWriter, request *http.Request){
		res.processRequest(request)
		body, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		writer.Write([]byte(body))
	})
	return res
}

func (ctx ServerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx.handler(writer, request)
}

func (ctx ServerHandler) Start() {
	server := &http.Server{Handler:ctx}
	server.Serve(ctx.listener)
}

func (ctx ServerHandler) Close() {
	ctx.listener.Close()
}

func (ctx * ServerHandler) processRequest(request * http.Request) {
	body := ""
	if request.Body != nil {
		data , _ := ioutil.ReadAll(request.Body)
		if data != nil {
			body = string(data)
		}
	}
	ctx.Body = string(body)
	ctx.Headers = toLocalHeaders(request.Header)
	delete(ctx.Headers, "Accept-Encoding")
	delete(ctx.Headers, "User-Agent")
}
