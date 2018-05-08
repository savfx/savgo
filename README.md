## go工具库

### 数据类型转换

- Ptr函数尝试转换,失败则为nil
  - convert.BoolPtr(value interface{}) *bool
    - 数字类型不为0则为true
    - 字符串(字节数组)值转为false
    - 字符串(字节数组)值为"1", "true", "TRUE", "True", "on"转为true
    - 字符串(字节数组)值为"null"转为nil
    - 其他均为false
  - convert.StringPtr(value interface{}) *string
    - 布尔类型转为"true"或"false"
    - 整类型转为十进制字符串
    - 浮点类型转为64位字符串
  - convert.Float32Ptr(value interface{}) *float32
    - 布尔类型转为0或1
    - 数字类型转为64位浮点数,再强制转为32位浮点数(越界?)
    - 字符串(字节数组)值当作十进制数字处理
  - convert.Float64Ptr(value interface{}) *float64
  - convert.IntPtr(value interface{}) *int
    - 布尔类型转为0或1
    - 浮点数转为64位整形后降级, 越界则为nil
    - 字符串(字节数组)值当作十进制数字处理, 越界则为nil
  - convert.UintPtr(value interface{}) *uint
  - convert.Int16Ptr(value interface{}) *int16
  - convert.Int32Ptr(value interface{}) *int32
  - convert.Int64Ptr(value interface{}) *int64
  - convert.Int8Ptr(value interface{}) *int8
  - convert.Uint16Ptr(value interface{}) *uint16
  - convert.Uint32Ptr(value interface{}) *uint32
  - convert.Uint64Ptr(value interface{}) *uint64
  - convert.Uint8Ptr(value interface{}) *uint8
- 若Ptr转换为nil则使用默认值
  - convert.BoolVal(value interface{}) (res bool)
  - convert.StringVal(value interface{}) (res string)
  - convert.Float32Val(value interface{}) (res float32)
  - convert.Float64Val(value interface{}) (res float64)
  - convert.Int16Val(value interface{}) (res int16)
  - convert.Int32Val(value interface{}) (res int32)
  - convert.Int64Val(value interface{}) (res int64)
  - convert.Int8Val(value interface{}) (res int8)
  - convert.IntVal(value interface{}) (res int)
  - convert.Uint16Val(value interface{}) (res uint16)
  - convert.Uint32Val(value interface{}) (res uint32)
  - convert.Uint64Val(value interface{}) (res uint64)
  - convert.Uint8Val(value interface{}) (res uint8)
  - convert.UintVal(value interface{}) (res uint)

```go

package main

import (
  "github.com/savfx/savgo/util/convert"
  "fmt"
)

func main() {
  fmt.Println("1" == *convert.StringPtr(1))
  fmt.Println(true == *convert.BoolPtr(1))
  fmt.Println(1 == *convert.IntPtr("1"))
  fmt.Println(1 == *convert.Float32Ptr("1"))

  fmt.Println("" == convert.StringVal(nil))
  fmt.Println(false == convert.BoolVal(nil))
  fmt.Println(0 == convert.IntVal(nil))
  fmt.Println(0 == convert.Float32Val(nil))
}

```

### 数据读取

- ValueAccess值读取
  - convert.NewValueAccess(data interface{}) *ValueAccess
  - func (self *ValueAccess) Set(data interface{}) *ValueAccess
  - func (self ValueAccess) IsArray() bool
  - func (self *ValueAccess) Object() *ObjectAccess
  - func (self *ValueAccess) Array() *ArrayAccess
  - func (self ValueAccess) IsObject() bool
  - func (self ValueAccess) Bool() bool
  - func (self ValueAccess) BoolPtr() *bool
  - func (self ValueAccess) Float32() float32
  - func (self ValueAccess) Float32Ptr() *float32
  - func (self ValueAccess) Float64() float64
  - func (self ValueAccess) Float64Ptr() *float64
  - func (self ValueAccess) Int() int
  - func (self ValueAccess) Int16() int16
  - func (self ValueAccess) Int16Ptr() *int16
  - func (self ValueAccess) Int32() int32
  - func (self ValueAccess) Int32Ptr() *int32
  - func (self ValueAccess) Int64() int64
  - func (self ValueAccess) Int64Ptr() *int64
  - func (self ValueAccess) Int8() int8
  - func (self ValueAccess) Int8Ptr() *int8
  - func (self ValueAccess) IntPtr() *int
  - func (self ValueAccess) String() string
  - func (self ValueAccess) StringPtr() *string
  - func (self ValueAccess) Uint() uint
  - func (self ValueAccess) Uint16() uint16
  - func (self ValueAccess) Uint16Ptr() *uint16
  - func (self ValueAccess) Uint32() uint32
  - func (self ValueAccess) Uint32Ptr() *uint32
  - func (self ValueAccess) Uint64() uint64
  - func (self ValueAccess) Uint64Ptr() *uint64
  - func (self ValueAccess) Uint8() uint8
  - func (self ValueAccess) Uint8Ptr() *uint8
  - func (self ValueAccess) UintPtr() *uint
- ArrayAccess数组读取
  - convert.NewArrayAccess(data []interface{}) *ArrayAccess
  - func (self ArrayAccess) Field(index int) *ValueAccess
  - func (self ArrayAccess) ForEach(callback func(int, *ValueAccess) bool)
  - func (self ArrayAccess) Len() int
  - func (self *ArrayAccess) Set(data []interface{}) *ArrayAccess
- ObjectAccess对象读取
  - convert.NewObjectAccess(data map[string]interface{}) *ObjectAccess
  - func (self *ObjectAccess) Set(data map[string]interface{}) *ObjectAccess
  - func (self ObjectAccess) Has(name string) bool
  - func (self ObjectAccess) HasArray(name string) bool
  - func (self ObjectAccess) HasObject(name string) bool
  - func (self ObjectAccess) GetRaw(name string) interface{}
  - func (self ObjectAccess) GetValue(name string) *ValueAccess
  - func (self ObjectAccess) GetObject(name string) *ObjectAccess
  - func (self ObjectAccess) GetArray(name string) *ArrayAccess
  - func (self ObjectAccess) ForEach(callback func(string, *ValueAccess) bool)
  - func (self ObjectAccess) GetBool(name string) bool
  - func (self ObjectAccess) GetBoolPtr(name string) *bool
  - func (self ObjectAccess) GetFloat32(name string) float32
  - func (self ObjectAccess) GetFloat32Ptr(name string) *float32
  - func (self ObjectAccess) GetFloat64(name string) float64
  - func (self ObjectAccess) GetFloat64Ptr(name string) *float64
  - func (self ObjectAccess) GetInt(name string) int
  - func (self ObjectAccess) GetInt16(name string) int16
  - func (self ObjectAccess) GetInt16Ptr(name string) *int16
  - func (self ObjectAccess) GetInt32(name string) int32
  - func (self ObjectAccess) GetInt32Ptr(name string) *int32
  - func (self ObjectAccess) GetInt64(name string) int64
  - func (self ObjectAccess) GetInt64Ptr(name string) *int64
  - func (self ObjectAccess) GetInt8(name string) int8
  - func (self ObjectAccess) GetInt8Ptr(name string) *int8
  - func (self ObjectAccess) GetIntPtr(name string) *int
  - func (self ObjectAccess) GetString(name string) string
  - func (self ObjectAccess) GetStringPtr(name string) *string
  - func (self ObjectAccess) GetUint(name string) uint
  - func (self ObjectAccess) GetUint16(name string) uint16
  - func (self ObjectAccess) GetUint16Ptr(name string) *uint16
  - func (self ObjectAccess) GetUint32(name string) uint32
  - func (self ObjectAccess) GetUint32Ptr(name string) *uint32
  - func (self ObjectAccess) GetUint64(name string) uint64
  - func (self ObjectAccess) GetUint64Ptr(name string) *uint64
  - func (self ObjectAccess) GetUint8(name string) uint8
  - func (self ObjectAccess) GetUint8Ptr(name string) *uint8
  - func (self ObjectAccess) GetUintPtr(name string) *uint

```go

package main

import (
  "github.com/savfx/savgo/util/convert"
  "encoding/json"
  "fmt"
)

func main() {
  text := `{
    "name": "jetiny",
    "age": "35",
    "man": "true",
    "profile": {
       "company": "hfjy"
    },
    "followers": [
      {"name": "张三", "age": 10},
      {"name": "李四", "age": 20}
    ]
  }`
  jsonText := map[string]interface{}{}
  json.Unmarshal([]byte(text), &jsonText)
  object := convert.NewObjectAccess(jsonText)

  fmt.Println(object.Has("profile"))
  fmt.Println(object.HasObject("profile"))
  fmt.Println(object.GetObject("profile").GetString("company"))
  fmt.Println(object.Has("nofound"))
  fmt.Println(object.HasObject("nofound"))

  fmt.Println(object.Has("followers"))
  fmt.Println(object.HasArray("followers"))
  fmt.Println(object.GetValue("followers").IsArray())
  fmt.Println(object.GetArray("followers").Len())
  fmt.Println(object.GetArray("followers").Field(0).Object().GetString("name"))

  object.ForEach(func (name string, value * convert.ValueAccess ) bool {
    return false
  })
}
 
```

### 命名规则转换

- 四种命名规则互转
  - strcase.Camel(str string) string
  - strcase.Pascal(str string) string
  - strcase.Snake(str string) string
  - strcase.Hyphen(str string) string
- 辅助方法
  - strcase.Convert(str string, caseType CaseType) string
  - strcase.ConvertStringType(str string, caseTypeStr string) string

```go

package main

import (
  "github.com/savgo/util/strcase"
  "fmt"
)

func main() {
  fmt.Println(strcase.Camel("hello-world")) // "helloWorld"
  fmt.Println(strcase.Pascal("hello-world")) // "HelloWorld"
  fmt.Println(strcase.Snake("helloWorld")) // "hello_world"
  fmt.Println(strcase.Hyphen("helloWorld")) // "hello-world"

  fmt.Println(strcase.Convert("hello-world", strcase.CamelType)) // "helloWorld"
  fmt.Println(strcase.ConvertStringType("hello-world", "camel")) // "helloWorld"
}

```

### 密码散列

- 创建和校验
  - passwd.Create(plaintext string) (encoded string, err error)
  - passwd.Verify(plaintext, encoded string) (isValid bool, err error)
- `passwd.Create`
  - 类似php的`password_hash`
  - 目前只提供bcrypt算法
- `passwd.Verify`
  - 类似php的`password_verify`
- `passwd.IsCurrent`
  - 类似php的`password_needs_rehash`
  - 由于只有一个bcrypt所以暂时没用

```go

package main

import (
  "github.com/savfx/savgo/util/passwd"
  "fmt"
)

func main() {
  plaintext := "password"
  pwd, _ := passwd.Create(plaintext)
  isValid, _ := passwd.Verify(plaintext, pwd)
  fmt.Println(pwd) // $2a$08$...
  fmt.Println(len(pwd)) // 60
  fmt.Println(isValid) // true
}

```
