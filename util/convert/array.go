package convert

type ArrayAccess struct {
  data []interface{}
  valueAccess ValueAccess
}

func NewArrayAccess(data []interface{}) (*ArrayAccess) {
  return &ArrayAccess{data, ValueAccess{}}
}

func (self ArrayAccess) Len() int {
  return len(self.data)
}

func (self * ArrayAccess) Set(data []interface{}) *ArrayAccess {
  self.data = data
  return self
}

func (self ArrayAccess) Field(index int) (*ValueAccess) {
  if index < 0 || index >= len(self.data) {
    return self.valueAccess.Set(nil)
  }
  return self.valueAccess.Set(self.data[index])
}

func (self ArrayAccess) ForEach(callback func( int, *ValueAccess)(bool)) {
  for index, val := range self.data {
    if !callback(index, self.valueAccess.Set(val)) {
      return
    }
  }
}
