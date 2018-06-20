package convert

import (
	"strconv"
	"strings"
)

type Values map[string][]string

func (v Values) GetArrayLen(prefix string) (bool, int) {
	if vs, ok := v[prefix+"[]"]; ok {
		return true, len(vs)
	}
	n := 0
	prefix += "["
	for key, _ := range v {
		if pos := strings.Index(key, prefix); pos != -1 {
			str := strings.TrimLeft(key, prefix)
			if pos2 := strings.Index(str, "]"); pos2 != -1 {
				index := str[0:pos2]
				if ret, err := strconv.ParseInt(index, 10, 0); err == nil {
					res := int(ret) + 1
					if n < res {
						n = res
					}
				}
			}
		}
	}
	return false, n
}

type FormObject struct {
	values      *Values
	prefix      string
	valueAccess *ValueAccess
}

func NewFormObject(values map[string][]string) *FormObject {
	val := Values(values)
	ret := &FormObject{}
	ret.values = &val
	return ret
}

func createFormObject(values *Values, prefix string) *FormObject {
	ret := &FormObject{}
	ret.values = values
	ret.prefix = prefix
	return ret
}

func (ctx FormObject) Prefix(key string) string {
	if len(ctx.prefix) == 0 {
		return key
	}
	return ctx.prefix + "[" + key + "]"
}

func (ctx *FormObject) SetPrefix(key string) *FormObject {
	ctx.prefix = key
	return ctx
}

func (ctx *FormObject) GetObject(key string) *FormObject {
	return createFormObject(ctx.values, ctx.Prefix(key))
}

func (ctx *FormObject) GetArray(key string) *FormArray {
	return createFormArray(ctx.values, ctx.Prefix(key))
}

func (ctx FormObject) GetRaw(name string) interface{} {
	if vs, ok := (*ctx.values)[ctx.Prefix(name)]; ok {
		if len(vs) > 0 {
			return vs[0]
		}
	}
	return nil
}

func (ctx FormObject) GetValue(name string) *ValueAccess {
	if ctx.valueAccess == nil {
		ctx.valueAccess = NewValueAccess(nil)
	}
	return ctx.valueAccess.Set(ctx.GetRaw(name))
}

func (ctx FormObject) GetStringPtr(name string) *string {
	return StringPtr(ctx.GetRaw(name))
}

func (ctx FormObject) GetBoolPtr(name string) *bool {
	return BoolPtr(ctx.GetRaw(name))
}

func (ctx FormObject) GetIntPtr(name string) *int {
	return IntPtr(ctx.GetRaw(name))
}

func (ctx FormObject) GetUintPtr(name string) *uint {
	return UintPtr(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt8Ptr(name string) *int8 {
	return Int8Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint8Ptr(name string) *uint8 {
	return Uint8Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt16Ptr(name string) *int16 {
	return Int16Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint16Ptr(name string) *uint16 {
	return Uint16Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt32Ptr(name string) *int32 {
	return Int32Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint32Ptr(name string) *uint32 {
	return Uint32Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt64Ptr(name string) *int64 {
	return Int64Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint64Ptr(name string) *uint64 {
	return Uint64Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetFloat32Ptr(name string) *float32 {
	return Float32Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetFloat64Ptr(name string) *float64 {
	return Float64Ptr(ctx.GetRaw(name))
}

func (ctx FormObject) GetString(name string) string {
	return StringVal(ctx.GetRaw(name))
}

func (ctx FormObject) GetBool(name string) bool {
	return BoolVal(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt(name string) int {
	return IntVal(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint(name string) uint {
	return UintVal(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt8(name string) int8 {
	return Int8Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint8(name string) uint8 {
	return Uint8Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt16(name string) int16 {
	return Int16Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint16(name string) uint16 {
	return Uint16Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt32(name string) int32 {
	return Int32Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint32(name string) uint32 {
	return Uint32Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetInt64(name string) int64 {
	return Int64Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetUint64(name string) uint64 {
	return Uint64Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetFloat32(name string) float32 {
	return Float32Val(ctx.GetRaw(name))
}

func (ctx FormObject) GetFloat64(name string) float64 {
	return Float64Val(ctx.GetRaw(name))
}

type FormArray struct {
	values     *Values
	prefix     string
	single     bool
	length     int
	field      *FormObject
	value      *ValueAccess
	arrayValue []string
}

func createFormArray(values *Values, prefix string) *FormArray {
	ret := &FormArray{}
	ret.values = values
	ret.prefix = prefix
	ret.single, ret.length = values.GetArrayLen(prefix)
	if ret.single {
		key := prefix + "[]"
		ret.arrayValue = (*values)[key]
		ret.value = &ValueAccess{}
	} else {
		ret.field = &FormObject{}
		ret.field.values = values
	}
	return ret
}

func (ctx *FormArray) Len() int {
	return ctx.length
}

func (ctx *FormArray) Single() bool {
	return ctx.single
}

func (ctx FormArray) Prefix(key string) string {
	if len(ctx.prefix) == 0 {
		return key
	}
	return ctx.prefix + "[" + key + "]"
}

func (ctx FormArray) Field(index int) *FormObject {
	if index < 0 || index >= ctx.length {
		return nil
	}
	return ctx.field.SetPrefix(ctx.Prefix(strconv.Itoa(index)))
}

func (ctx FormArray) EachField(callback func(int, *FormObject)) {
	for i := 0; i < ctx.length; i++ {
		callback(i, ctx.field.SetPrefix(ctx.Prefix(strconv.Itoa(i))))
	}
}

func (ctx FormArray) Value(index int) *ValueAccess {
	if index < 0 || index >= ctx.length {
		return ctx.value.Set(nil)
	}
	return ctx.value.Set(ctx.arrayValue[index])
}

func (ctx FormArray) EachValue(callback func(int, *ValueAccess)) {
	for i := 0; i < ctx.length; i++ {
		callback(i, ctx.value.Set(ctx.arrayValue[i]))
	}
}
