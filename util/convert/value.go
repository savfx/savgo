package convert

type ValueAccess struct {
	data   interface{}
	array  *ArrayAccess
	object *ObjectAccess
}

func NewValueAccess(data interface{}) (*ValueAccess) {
	res := &ValueAccess{}
	res.data = data
	return res
}

func (ctx *ValueAccess) Set(data interface{}) (*ValueAccess) {
	ctx.data = data
	return ctx
}

func (ctx ValueAccess) IsArray() bool {
	switch ctx.data.(type) {
	case []interface{}:
		return true
	}
	return false
}

func (ctx *ValueAccess) Array() *ArrayAccess {
	if ctx.array == nil {
		ctx.array = NewArrayAccess(nil)
	}
	switch val := ctx.data.(type) {
	case []interface{}:
		ctx.array.Set(val)
	}
	return ctx.array
}

func (ctx ValueAccess) IsObject() bool {
	switch ctx.data.(type) {
	case map[string]interface{}:
		return true
	}
	return false
}

func (ctx *ValueAccess) Object() *ObjectAccess {
	if ctx.object == nil {
		ctx.object = NewObjectAccess(nil)
	}
	switch val := ctx.data.(type) {
	case map[string]interface{}:
		ctx.object.Set(val)
	}
	return ctx.object
}

func (ctx ValueAccess) StringPtr() *string {
	return StringPtr(ctx.data)
}

func (ctx ValueAccess) BoolPtr() *bool {
	return BoolPtr(ctx.data)
}

func (ctx ValueAccess) IntPtr() *int {
	return IntPtr(ctx.data)
}

func (ctx ValueAccess) UintPtr() *uint {
	return UintPtr(ctx.data)
}

func (ctx ValueAccess) Int8Ptr() *int8 {
	return Int8Ptr(ctx.data)
}

func (ctx ValueAccess) Uint8Ptr() *uint8 {
	return Uint8Ptr(ctx.data)
}

func (ctx ValueAccess) Int16Ptr() *int16 {
	return Int16Ptr(ctx.data)
}

func (ctx ValueAccess) Uint16Ptr() *uint16 {
	return Uint16Ptr(ctx.data)
}

func (ctx ValueAccess) Int32Ptr() *int32 {
	return Int32Ptr(ctx.data)
}

func (ctx ValueAccess) Uint32Ptr() *uint32 {
	return Uint32Ptr(ctx.data)
}

func (ctx ValueAccess) Int64Ptr() *int64 {
	return Int64Ptr(ctx.data)
}

func (ctx ValueAccess) Uint64Ptr() *uint64 {
	return Uint64Ptr(ctx.data)
}

func (ctx ValueAccess) Float32Ptr() *float32 {
	return Float32Ptr(ctx.data)
}

func (ctx ValueAccess) Float64Ptr() *float64 {
	return Float64Ptr(ctx.data)
}

func (ctx ValueAccess) String() string {
	return StringVal(ctx.data)
}

func (ctx ValueAccess) Bool() bool {
	return BoolVal(ctx.data)
}

func (ctx ValueAccess) Int() int {
	return IntVal(ctx.data)
}

func (ctx ValueAccess) Uint() uint {
	return UintVal(ctx.data)
}

func (ctx ValueAccess) Int8() int8 {
	return Int8Val(ctx.data)
}

func (ctx ValueAccess) Uint8() uint8 {
	return Uint8Val(ctx.data)
}

func (ctx ValueAccess) Int16() int16 {
	return Int16Val(ctx.data)
}

func (ctx ValueAccess) Uint16() uint16 {
	return Uint16Val(ctx.data)
}

func (ctx ValueAccess) Int32() int32 {
	return Int32Val(ctx.data)
}

func (ctx ValueAccess) Uint32() uint32 {
	return Uint32Val(ctx.data)
}

func (ctx ValueAccess) Int64() int64 {
	return Int64Val(ctx.data)
}

func (ctx ValueAccess) Uint64() uint64 {
	return Uint64Val(ctx.data)
}

func (ctx ValueAccess) Float32() float32 {
	return Float32Val(ctx.data)
}

func (ctx ValueAccess) Float64() float64 {
	return Float64Val(ctx.data)
}
