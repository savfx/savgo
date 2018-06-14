package convert

type ArrayAccess struct {
	data        []interface{}
	valueAccess ValueAccess
}

func NewArrayAccess(data []interface{}) (*ArrayAccess) {
	return &ArrayAccess{data, ValueAccess{}}
}

func (ctx ArrayAccess) Len() int {
	return len(ctx.data)
}

func (ctx *ArrayAccess) Set(data []interface{}) *ArrayAccess {
	ctx.data = data
	return ctx
}

func (ctx ArrayAccess) Field(index int) (*ValueAccess) {
	if index < 0 || index >= len(ctx.data) {
		return ctx.valueAccess.Set(nil)
	}
	return ctx.valueAccess.Set(ctx.data[index])
}

func (ctx ArrayAccess) ForEach(callback func(int, *ValueAccess) (bool)) {
	for index, val := range ctx.data {
		if !callback(index, ctx.valueAccess.Set(val)) {
			return
		}
	}
}
