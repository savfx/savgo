package sav

type BaseRequest struct {
	Headers    map[string]string
	DataSource *BaseDataSource
}

func (ctx BaseRequest) GetHeaders() map[string]string {
	return ctx.Headers
}

func (ctx BaseRequest) GetData() DataSource {
	return ctx.DataSource
}

type BaseResponse struct {
	StatusCode int
	Headers    map[string]string
	DataSource *BaseDataSource
	Body string
}

func (ctx BaseResponse) GetStatusCode() int {
	return ctx.StatusCode
}

func (ctx BaseResponse) GetHeaders() map[string]string {
	return ctx.Headers
}

func (ctx BaseResponse) GetData() DataSource {
	return ctx.DataSource
}
