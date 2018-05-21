package convert

import (
	"strconv"
	"strings"
)

type Values map[string][]string

func (v Values) GetArrayLen (prefix string) (bool, int) {
	if vs, ok := v[prefix + "[]"]; ok {
		return true, len(vs)
	}
	n := 0
	prefix += "["
	for key, _ := range v {
		if pos := strings.Index(key, prefix); pos != -1 {
			str := strings.TrimLeft(key, prefix)
			if pos2 := strings.Index(str, "]"); pos2 != -1 {
				index := str[0: pos2]
				if ret, err := strconv.ParseInt(index, 10, 0); err == nil {
					nret := int(ret) + 1
					if n < nret {
						n = nret
					}
				}
			}
		}
	}
	return false, n
}

type FormObject struct {
	values * Values
	prefix string
	valueAccess * ValueAccess
}

func NewFormObject(values map[string][]string) * FormObject {
	vals := Values(values)
	ret := &FormObject{}
	ret.values = &vals
	return ret
}

func createFormObject(values * Values, prefix string) * FormObject {
	ret := &FormObject{}
	ret.values = values
	ret.prefix = prefix
	return ret
}

func (self FormObject) Prefix(key string) string {
	if len(self.prefix) == 0 {
		return key
	}
	return self.prefix + "[" + key + "]"
}

func (self * FormObject) SetPrefix(key string) *FormObject {
	self.prefix = key
	return self
}

func (self * FormObject) GetObject(key string) * FormObject{
	return createFormObject(self.values, self.Prefix(key))
}

func (self * FormObject) GetArray(key string) * FormArray{
	return createFormArray(self.values, self.Prefix(key))
}

func (self FormObject) GetRaw(name string) interface{} {
	if vs, ok := (*self.values)[self.Prefix(name)]; ok {
		if len(vs) > 0 {
			return vs[0]
		}
	}
	return nil
}

func (self FormObject) GetValue(name string) *ValueAccess {
  if self.valueAccess == nil {
    self.valueAccess = NewValueAccess(nil)
  }
  return self.valueAccess.Set(self.GetRaw(name))
}

func (self FormObject) GetStringPtr(name string) *string {
  return StringPtr(self.GetRaw(name))
}

func (self FormObject) GetBoolPtr(name string) *bool {
  return BoolPtr(self.GetRaw(name))
}

func (self FormObject) GetIntPtr(name string) *int {
  return IntPtr(self.GetRaw(name))
}

func (self FormObject) GetUintPtr(name string) *uint {
  return UintPtr(self.GetRaw(name))
}

func (self FormObject) GetInt8Ptr(name string) *int8 {
  return Int8Ptr(self.GetRaw(name))
}

func (self FormObject) GetUint8Ptr(name string) *uint8 {
  return Uint8Ptr(self.GetRaw(name))
}

func (self FormObject) GetInt16Ptr(name string) *int16 {
  return Int16Ptr(self.GetRaw(name))
}

func (self FormObject) GetUint16Ptr(name string) *uint16 {
  return Uint16Ptr(self.GetRaw(name))
}

func (self FormObject) GetInt32Ptr(name string) *int32 {
  return Int32Ptr(self.GetRaw(name))
}

func (self FormObject) GetUint32Ptr(name string) *uint32 {
  return Uint32Ptr(self.GetRaw(name))
}

func (self FormObject) GetInt64Ptr(name string) *int64 {
  return Int64Ptr(self.GetRaw(name))
}

func (self FormObject) GetUint64Ptr(name string) *uint64 {
  return Uint64Ptr(self.GetRaw(name))
}

func (self FormObject) GetFloat32Ptr(name string) *float32 {
  return Float32Ptr(self.GetRaw(name))
}

func (self FormObject) GetFloat64Ptr(name string) *float64 {
  return Float64Ptr(self.GetRaw(name))
}

func (self FormObject) GetString(name string) string {
  return StringVal(self.GetRaw(name))
}

func (self FormObject) GetBool(name string) bool {
  return BoolVal(self.GetRaw(name))
}

func (self FormObject) GetInt(name string) int {
  return IntVal(self.GetRaw(name))
}

func (self FormObject) GetUint(name string) uint {
  return UintVal(self.GetRaw(name))
}

func (self FormObject) GetInt8(name string) int8 {
  return Int8Val(self.GetRaw(name))
}

func (self FormObject) GetUint8(name string) uint8 {
  return Uint8Val(self.GetRaw(name))
}

func (self FormObject) GetInt16(name string) int16 {
  return Int16Val(self.GetRaw(name))
}

func (self FormObject) GetUint16(name string) uint16 {
  return Uint16Val(self.GetRaw(name))
}

func (self FormObject) GetInt32(name string) int32 {
  return Int32Val(self.GetRaw(name))
}

func (self FormObject) GetUint32(name string) uint32 {
  return Uint32Val(self.GetRaw(name))
}

func (self FormObject) GetInt64(name string) int64 {
  return Int64Val(self.GetRaw(name))
}

func (self FormObject) GetUint64(name string) uint64 {
  return Uint64Val(self.GetRaw(name))
}

func (self FormObject) GetFloat32(name string) float32 {
  return Float32Val(self.GetRaw(name))
}

func (self FormObject) GetFloat64(name string) float64 {
  return Float64Val(self.GetRaw(name))
}

type FormArray struct {
	values * Values
	prefix string
	single bool
	length int
	field* FormObject
	value* ValueAccess
	arrayValue []string
}

func createFormArray(values * Values, prefix string) * FormArray{
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

func (self * FormArray) Len() int {
	return self.length
}

func (self * FormArray) Single() bool {
	return self.single
}

func (self FormArray) Prefix(key string) string {
	if len(self.prefix) == 0 {
		return key
	}
	return self.prefix + "[" + key + "]"
}

func (self FormArray) Field(index int) (*FormObject) {
  if index < 0 || index >= self.length {
    return nil
  }
  return self.field.SetPrefix(self.Prefix(strconv.Itoa(index)))
}

func (self FormArray) EachField(callback func( int, *FormObject)()) {
	for i :=0 ; i < self.length; i++ {
		callback(i, self.field.SetPrefix(self.Prefix(strconv.Itoa(i))))
	}
}

func (self FormArray) Value(index int) (* ValueAccess) {
  if index < 0 || index >= self.length {
    return self.value.Set(nil)
  }
  return self.value.Set(self.arrayValue[index])
}

func (self FormArray) EachValue(callback func( int, * ValueAccess)()) {
	for i :=0 ; i < self.length; i++ {
		callback(i, self.value.Set(self.arrayValue[i]))
	}
}
