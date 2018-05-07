package convert

type ValueAccess struct {
  data interface{}
  array *ArrayAccess
  object *ObjectAccess
}

func NewValueAccess(data interface{}) (*ValueAccess) {
  res := &ValueAccess{}
  res.data = data
  return res
}

func (self * ValueAccess) Set(data interface{}) (*ValueAccess){
  self.data = data
  return self
}

func (self ValueAccess) IsArray() bool {
  switch self.data.(type) {
    case []interface{}:
      return true
  }
  return false
}

func (self * ValueAccess) Array() *ArrayAccess {
  if (self.array == nil) {
    self.array = NewArrayAccess(nil)
  }
  switch val := self.data.(type) {
    case []interface{}:
      self.array.Set(val)
  }
  return self.array
}

func (self ValueAccess) IsObject() bool {
  switch self.data.(type) {
    case map[string]interface{}:
      return true
  }
  return false
}

func (self * ValueAccess) Object() *ObjectAccess {
  if (self.object == nil) {
    self.object = NewObjectAccess(nil)
  }
  switch val := self.data.(type) {
    case map[string]interface{}:
      self.object.Set(val)
  }
  return self.object
}

func (self ValueAccess) StringPtr() *string {
  return StringPtr(self.data)
}

func (self ValueAccess) BoolPtr() *bool {
  return BoolPtr(self.data)
}

func (self ValueAccess) IntPtr() *int {
  return IntPtr(self.data)
}

func (self ValueAccess) UintPtr() *uint {
  return UintPtr(self.data)
}

func (self ValueAccess) Int8Ptr() *int8 {
  return Int8Ptr(self.data)
}

func (self ValueAccess) Uint8Ptr() *uint8 {
  return Uint8Ptr(self.data)
}

func (self ValueAccess) Int16Ptr() *int16 {
  return Int16Ptr(self.data)
}

func (self ValueAccess) Uint16Ptr() *uint16 {
  return Uint16Ptr(self.data)
}

func (self ValueAccess) Int32Ptr() *int32 {
  return Int32Ptr(self.data)
}

func (self ValueAccess) Uint32Ptr() *uint32 {
  return Uint32Ptr(self.data)
}

func (self ValueAccess) Int64Ptr() *int64 {
  return Int64Ptr(self.data)
}

func (self ValueAccess) Uint64Ptr() *uint64 {
  return Uint64Ptr(self.data)
}

func (self ValueAccess) Float32Ptr() *float32 {
  return Float32Ptr(self.data)
}

func (self ValueAccess) Float64Ptr() *float64 {
  return Float64Ptr(self.data)
}

func (self ValueAccess) String() string {
  return StringVal(self.data)
}

func (self ValueAccess) Bool() bool {
  return BoolVal(self.data)
}

func (self ValueAccess) Int() int {
  return IntVal(self.data)
}

func (self ValueAccess) Uint() uint {
  return UintVal(self.data)
}

func (self ValueAccess) Int8() int8 {
  return Int8Val(self.data)
}

func (self ValueAccess) Uint8() uint8 {
  return Uint8Val(self.data)
}

func (self ValueAccess) Int16() int16 {
  return Int16Val(self.data)
}

func (self ValueAccess) Uint16() uint16 {
  return Uint16Val(self.data)
}

func (self ValueAccess) Int32() int32 {
  return Int32Val(self.data)
}

func (self ValueAccess) Uint32() uint32 {
  return Uint32Val(self.data)
}

func (self ValueAccess) Int64() int64 {
  return Int64Val(self.data)
}

func (self ValueAccess) Uint64() uint64 {
  return Uint64Val(self.data)
}

func (self ValueAccess) Float32() float32 {
  return Float32Val(self.data)
}

func (self ValueAccess) Float64() float64 {
  return Float64Val(self.data)
}
