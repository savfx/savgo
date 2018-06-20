package convert

type ObjectAccess struct {
	data        map[string]interface{}
	valueAccess *ValueAccess
}

func NewObjectAccess(data map[string]interface{}) *ObjectAccess {
	res := &ObjectAccess{}
	return res.Set(data)
}

func (ctx *ObjectAccess) Set(data map[string]interface{}) *ObjectAccess {
	if data == nil {
		ctx.data = map[string]interface{}{}
	} else {
		ctx.data = data
	}
	return ctx
}

func (ctx ObjectAccess) ForEach(callback func(string, *ValueAccess) bool) {
	if ctx.valueAccess == nil {
		ctx.valueAccess = NewValueAccess(nil)
	}
	for index, val := range ctx.data {
		if !callback(index, ctx.valueAccess.Set(val)) {
			return
		}
	}
}

func (ctx ObjectAccess) Has(name string) bool {
	if _, ok := ctx.data[name]; ok {
		return true
	}
	return false
}

func (ctx ObjectAccess) GetValue(name string) *ValueAccess {
	if ctx.valueAccess == nil {
		ctx.valueAccess = NewValueAccess(nil)
	}
	return ctx.valueAccess.Set(ctx.data[name])
}

func (ctx ObjectAccess) HasObject(name string) bool {
	if value, ok := ctx.data[name]; ok {
		switch value.(type) {
		case map[string]interface{}:
			return true
		}
	}
	return false
}

func (ctx ObjectAccess) GetObject(name string) *ObjectAccess {
	if ctx.HasObject(name) {
		return NewObjectAccess(ctx.GetRaw(name).(map[string]interface{}))
	}
	return NewObjectAccess(map[string]interface{}{})
}

func (ctx ObjectAccess) HasArray(name string) bool {
	if value, ok := ctx.data[name]; ok {
		switch value.(type) {
		case []interface{}:
			return true
		}
	}
	return false
}

func (ctx ObjectAccess) GetArray(name string) *ArrayAccess {
	if ctx.HasArray(name) {
		return NewArrayAccess(ctx.GetRaw(name).([]interface{}))
	}
	return NewArrayAccess([]interface{}{})
}

func (ctx ObjectAccess) GetRaw(name string) interface{} {
	return ctx.data[name]
}

func (ctx ObjectAccess) GetStringPtr(name string) *string {
	return StringPtr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetBoolPtr(name string) *bool {
	return BoolPtr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetIntPtr(name string) *int {
	return IntPtr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUintPtr(name string) *uint {
	return UintPtr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt8Ptr(name string) *int8 {
	return Int8Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint8Ptr(name string) *uint8 {
	return Uint8Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt16Ptr(name string) *int16 {
	return Int16Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint16Ptr(name string) *uint16 {
	return Uint16Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt32Ptr(name string) *int32 {
	return Int32Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint32Ptr(name string) *uint32 {
	return Uint32Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt64Ptr(name string) *int64 {
	return Int64Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint64Ptr(name string) *uint64 {
	return Uint64Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetFloat32Ptr(name string) *float32 {
	return Float32Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetFloat64Ptr(name string) *float64 {
	return Float64Ptr(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetString(name string) string {
	return StringVal(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetBool(name string) bool {
	return BoolVal(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt(name string) int {
	return IntVal(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint(name string) uint {
	return UintVal(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt8(name string) int8 {
	return Int8Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint8(name string) uint8 {
	return Uint8Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt16(name string) int16 {
	return Int16Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint16(name string) uint16 {
	return Uint16Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt32(name string) int32 {
	return Int32Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint32(name string) uint32 {
	return Uint32Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetInt64(name string) int64 {
	return Int64Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetUint64(name string) uint64 {
	return Uint64Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetFloat32(name string) float32 {
	return Float32Val(ctx.GetRaw(name))
}

func (ctx ObjectAccess) GetFloat64(name string) float64 {
	return Float64Val(ctx.GetRaw(name))
}
