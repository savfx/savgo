package convert

type IArrayDecoder interface {
  GetObject(index int) IObjectDecoder
  HasObject(index int) bool
}

type IObjectDecoder interface {
  GetStringPtr(name string) *string
  GetString(name string) string
  GetBoolPtr(name string) *bool
  GetBool(name string) bool
  GetIntPtr(name string) *int
  GetInt(name string) int
  GetUintPtr(name string) *uint
  GetUint(name string) uint
  GetInt8Ptr(name string) *int8
  GetInt8(name string) int8
  GetUint8Ptr(name string) *uint8
  GetUint8(name string) uint8
  GetInt16Ptr(name string) *int16
  GetInt16(name string) int16
  GetUint16Ptr(name string) *uint16
  GetUint16(name string) uint16
  GetInt32Ptr(name string) *int32
  GetInt32(name string) int32
  GetUint32Ptr(name string) *uint32
  GetUint32(name string) uint32
  GetInt64Ptr(name string) *int64
  GetInt64(name string) int64
  GetUint64Ptr(name string) *uint64
  GetUint64(name string) uint64
  GetFloat32Ptr(name string) *float32
  GetFloat32(name string) float32
  GetFloat64Ptr(name string) *float64
  GetFloat64(name string) float64
  GetObject(name string) IObjectDecoder
  HasObject(name string) bool
  // GetArray(name string) IArrayDecoder
  HasArray(name string) bool
  GetRaw(name string) interface{}
  Has(name string) bool
}

type ObjectWalker struct {
  data map[string]interface{}
}

func NewObjectWalker(data map[string]interface{}) (*ObjectWalker){
  res := &ObjectWalker{}
  res.data = data
  return res
}

func (self * ObjectWalker) GetRaw(name string) interface{} {
  return self.data[name]
}

func (self * ObjectWalker) Has(name string) bool {
  if _, ok := self.data[name]; ok {
    return true
  }
  return false
}

func (self * ObjectWalker) HasObject(name string) bool {
  if value, ok := self.data[name]; ok {
    switch value.(type) {
      case map[string]interface{}:
        return true
    }
  }
  return false
}

func (self * ObjectWalker) HasArray(name string) bool {
  if value, ok := self.data[name]; ok {
    switch value.(type) {
      case []interface{}:
        return true
    }
  }
  return false
}

type ObjectDecoder struct {
  walker *ObjectWalker
}

func NewObjectDecoder(value map[string]interface{}) (IObjectDecoder) {
  res := &ObjectDecoder{}
  res.walker = NewObjectWalker(value)
  return res
}

func (self ObjectDecoder) GetObject(name string) IObjectDecoder {
  if self.walker.HasObject(name) {
    return NewObjectDecoder(self.GetRaw(name).(map[string]interface{}))
  }
  return NewObjectDecoder(map[string]interface{}{})
}

func (self ObjectDecoder) HasObject(name string) bool {
  return self.walker.HasObject(name)
}

func (self ObjectDecoder) HasArray(name string) bool {
  return self.walker.HasArray(name)
}

func (self ObjectDecoder) GetRaw(name string) interface{} {
  return self.walker.GetRaw(name)
}

func (self ObjectDecoder) Has(name string) bool {
  return self.walker.Has(name)
}

func (self ObjectDecoder) GetStringPtr(name string) *string {
  return StringPtr(self.GetRaw(name))
}

func (self ObjectDecoder) GetBoolPtr(name string) *bool {
  return BoolPtr(self.GetRaw(name))
}

func (self ObjectDecoder) GetIntPtr(name string) *int {
  return IntPtr(self.GetRaw(name))
}

func (self ObjectDecoder) GetUintPtr(name string) *uint {
  return UintPtr(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt8Ptr(name string) *int8 {
  return Int8Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint8Ptr(name string) *uint8 {
  return Uint8Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt16Ptr(name string) *int16 {
  return Int16Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint16Ptr(name string) *uint16 {
  return Uint16Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt32Ptr(name string) *int32 {
  return Int32Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint32Ptr(name string) *uint32 {
  return Uint32Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt64Ptr(name string) *int64 {
  return Int64Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint64Ptr(name string) *uint64 {
  return Uint64Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetFloat32Ptr(name string) *float32 {
  return Float32Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetFloat64Ptr(name string) *float64 {
  return Float64Ptr(self.GetRaw(name))
}

func (self ObjectDecoder) GetString(name string) string {
  return StringVal(self.GetRaw(name))
}

func (self ObjectDecoder) GetBool(name string) bool {
  return BoolVal(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt(name string) int {
  return IntVal(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint(name string) uint {
  return UintVal(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt8(name string) int8 {
  return Int8Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint8(name string) uint8 {
  return Uint8Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt16(name string) int16 {
  return Int16Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint16(name string) uint16 {
  return Uint16Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt32(name string) int32 {
  return Int32Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint32(name string) uint32 {
  return Uint32Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetInt64(name string) int64 {
  return Int64Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetUint64(name string) uint64 {
  return Uint64Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetFloat32(name string) float32 {
  return Float32Val(self.GetRaw(name))
}

func (self ObjectDecoder) GetFloat64(name string) float64 {
  return Float64Val(self.GetRaw(name))
}
