package convert

type ObjectAccess struct {
  data map[string]interface{}
  valueAccess *ValueAccess
}

func NewObjectAccess(data map[string]interface{}) (*ObjectAccess) {
  res := &ObjectAccess{}
  return res.Set(data)
}

func (self * ObjectAccess) Set(data map[string]interface{}) *ObjectAccess {
  if data == nil {
    self.data = map[string]interface{}{}
  } else {
    self.data = data
  }
  return self
}

func (self ObjectAccess) ForEach(callback func(string, *ValueAccess)(bool)) {
  if self.valueAccess == nil {
    self.valueAccess = NewValueAccess(nil)
  }
  for index, val := range self.data {
    if !callback(index, self.valueAccess.Set(val)) {
      return
    }
  }
}

func (self ObjectAccess) Has(name string) bool {
  if _, ok := self.data[name]; ok {
    return true
  }
  return false
}

func (self ObjectAccess) GetValue(name string) *ValueAccess {
  if self.valueAccess == nil {
    self.valueAccess = NewValueAccess(nil)
  }
  return self.valueAccess.Set(self.data[name])
}

func (self ObjectAccess) HasObject(name string) bool {
  if value, ok := self.data[name]; ok {
    switch value.(type) {
      case map[string]interface{}:
        return true
    }
  }
  return false
}

func (self ObjectAccess) GetObject(name string) (*ObjectAccess) {
  if self.HasObject(name) {
    return NewObjectAccess(self.GetRaw(name).(map[string]interface{}))
  }
  return NewObjectAccess(map[string]interface{}{})
}

func (self ObjectAccess) HasArray(name string) bool {
  if value, ok := self.data[name]; ok {
    switch value.(type) {
      case []interface{}:
        return true
    }
  }
  return false
}

func (self ObjectAccess) GetArray(name string) (*ArrayAccess) {
  if self.HasArray(name) {
    return NewArrayAccess(self.GetRaw(name).([]interface{}))
  }
  return NewArrayAccess([]interface{}{})
}

func (self ObjectAccess) GetRaw(name string) interface{} {
  return self.data[name]
}

func (self ObjectAccess) GetStringPtr(name string) *string {
  return StringPtr(self.GetRaw(name))
}

func (self ObjectAccess) GetBoolPtr(name string) *bool {
  return BoolPtr(self.GetRaw(name))
}

func (self ObjectAccess) GetIntPtr(name string) *int {
  return IntPtr(self.GetRaw(name))
}

func (self ObjectAccess) GetUintPtr(name string) *uint {
  return UintPtr(self.GetRaw(name))
}

func (self ObjectAccess) GetInt8Ptr(name string) *int8 {
  return Int8Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetUint8Ptr(name string) *uint8 {
  return Uint8Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetInt16Ptr(name string) *int16 {
  return Int16Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetUint16Ptr(name string) *uint16 {
  return Uint16Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetInt32Ptr(name string) *int32 {
  return Int32Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetUint32Ptr(name string) *uint32 {
  return Uint32Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetInt64Ptr(name string) *int64 {
  return Int64Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetUint64Ptr(name string) *uint64 {
  return Uint64Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetFloat32Ptr(name string) *float32 {
  return Float32Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetFloat64Ptr(name string) *float64 {
  return Float64Ptr(self.GetRaw(name))
}

func (self ObjectAccess) GetString(name string) string {
  return StringVal(self.GetRaw(name))
}

func (self ObjectAccess) GetBool(name string) bool {
  return BoolVal(self.GetRaw(name))
}

func (self ObjectAccess) GetInt(name string) int {
  return IntVal(self.GetRaw(name))
}

func (self ObjectAccess) GetUint(name string) uint {
  return UintVal(self.GetRaw(name))
}

func (self ObjectAccess) GetInt8(name string) int8 {
  return Int8Val(self.GetRaw(name))
}

func (self ObjectAccess) GetUint8(name string) uint8 {
  return Uint8Val(self.GetRaw(name))
}

func (self ObjectAccess) GetInt16(name string) int16 {
  return Int16Val(self.GetRaw(name))
}

func (self ObjectAccess) GetUint16(name string) uint16 {
  return Uint16Val(self.GetRaw(name))
}

func (self ObjectAccess) GetInt32(name string) int32 {
  return Int32Val(self.GetRaw(name))
}

func (self ObjectAccess) GetUint32(name string) uint32 {
  return Uint32Val(self.GetRaw(name))
}

func (self ObjectAccess) GetInt64(name string) int64 {
  return Int64Val(self.GetRaw(name))
}

func (self ObjectAccess) GetUint64(name string) uint64 {
  return Uint64Val(self.GetRaw(name))
}

func (self ObjectAccess) GetFloat32(name string) float32 {
  return Float32Val(self.GetRaw(name))
}

func (self ObjectAccess) GetFloat64(name string) float64 {
  return Float64Val(self.GetRaw(name))
}
