package convert

import (
  "fmt"
  "github.com/a8m/expect"
  "testing"
  "encoding/json"
  "reflect"
)

func TestObjectDecoder(t *testing.T) {
  expect := expect.New(t)
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
  decoder := NewObjectDecoder(jsonText)
  expect(*decoder.GetStringPtr("name")).To.Equal("jetiny")
  expect(decoder.GetString("name")).To.Equal("jetiny")
  expect(*decoder.GetIntPtr("age")).To.Equal(35)
  expect(decoder.GetInt("age")).To.Equal(35)
  expect(*decoder.GetBoolPtr("man")).To.Equal(true)
  expect(decoder.GetBool("man")).To.Equal(true)
  expect(35 == *decoder.GetInt8Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetInt8("age")).To.Equal(true)
  expect(35 == *decoder.GetUint8Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetUint8("age")).To.Equal(true)
  expect(35 == *decoder.GetInt16Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetInt16("age")).To.Equal(true)
  expect(35 == *decoder.GetUint16Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetUint16("age")).To.Equal(true)
  expect(35 == *decoder.GetInt32Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetInt32("age")).To.Equal(true)
  expect(35 == *decoder.GetUint32Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetUint32("age")).To.Equal(true)
  expect(35 == *decoder.GetInt64Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetInt64("age")).To.Equal(true)
  expect(35 == *decoder.GetUint64Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetUint64("age")).To.Equal(true)
  expect(35 == *decoder.GetFloat32Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetFloat32("age")).To.Equal(true)
  expect(35 == *decoder.GetFloat64Ptr("age")).To.Equal(true)
  expect(35 == decoder.GetFloat64("age")).To.Equal(true)

  expect(decoder.GetString("nofont")).To.Equal("")
  expect(false == decoder.GetBool("nofont")).To.Equal(true)
  expect(0 == decoder.GetInt("nofont")).To.Equal(true)
  expect(0 == decoder.GetInt8("nofont")).To.Equal(true)
  expect(0 == decoder.GetInt16("nofont")).To.Equal(true)
  expect(0 == decoder.GetInt32("nofont")).To.Equal(true)
  expect(0 == decoder.GetInt64("nofont")).To.Equal(true)
  expect(0 == decoder.GetUint("nofont")).To.Equal(true)
  expect(0 == decoder.GetUint8("nofont")).To.Equal(true)
  expect(0 == decoder.GetUint16("nofont")).To.Equal(true)
  expect(0 == decoder.GetUint32("nofont")).To.Equal(true)
  expect(0 == decoder.GetUint64("nofont")).To.Equal(true)
  expect(0 == decoder.GetFloat32("nofont")).To.Equal(true)
  expect(0 == decoder.GetFloat64("nofont")).To.Equal(true)

  expect(decoder.Has("profile")).To.Equal(true)
  expect(decoder.HasObject("profile")).To.Equal(true)
  expect(decoder.GetObject("profile").GetString("company")).To.Equal("hfjy")
  expect(decoder.Has("nofound")).To.Equal(false)
  expect(decoder.HasObject("nofound")).To.Equal(false)

  expect(decoder.Has("followers")).To.Equal(true)
  expect(decoder.HasArray("followers")).To.Equal(true)
  fmt.Println()
}

func TestStringObjectDecoder(t *testing.T) {
  expect := expect.New(t)
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

  refType := reflect.TypeOf(jsonText["followers"])
  fmt.Println("----")
  fmt.Println(refType.String())
  fmt.Println(refType.Kind())
  // fmt.Println(refType.Len())
  // fmt.Println(refType.Key())
  // fmt.Println(refType.Elem())

  expect(true).To.Equal(true)

  // decoder := NewObjectDecoder(jsonText)
  // expect(*decoder.GetStringPtr("name")).To.Equal("jetiny")
  // expect(decoder.GetString("name")).To.Equal("jetiny")
  // expect(*decoder.GetIntPtr("age")).To.Equal(35)
  // expect(decoder.GetInt("age")).To.Equal(35)
  // expect(*decoder.GetBoolPtr("man")).To.Equal(true)
  // expect(decoder.GetBool("man")).To.Equal(true)
  // expect(35 == *decoder.GetInt8Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetInt8("age")).To.Equal(true)
  // expect(35 == *decoder.GetUint8Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetUint8("age")).To.Equal(true)
  // expect(35 == *decoder.GetInt16Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetInt16("age")).To.Equal(true)
  // expect(35 == *decoder.GetUint16Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetUint16("age")).To.Equal(true)
  // expect(35 == *decoder.GetInt32Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetInt32("age")).To.Equal(true)
  // expect(35 == *decoder.GetUint32Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetUint32("age")).To.Equal(true)
  // expect(35 == *decoder.GetInt64Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetInt64("age")).To.Equal(true)
  // expect(35 == *decoder.GetUint64Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetUint64("age")).To.Equal(true)
  // expect(35 == *decoder.GetFloat32Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetFloat32("age")).To.Equal(true)
  // expect(35 == *decoder.GetFloat64Ptr("age")).To.Equal(true)
  // expect(35 == decoder.GetFloat64("age")).To.Equal(true)

  // expect(decoder.GetString("nofont")).To.Equal("")
  // expect(false == decoder.GetBool("nofont")).To.Equal(true)
  // expect(0 == decoder.GetInt("nofont")).To.Equal(true)
  // expect(0 == decoder.GetInt8("nofont")).To.Equal(true)
  // expect(0 == decoder.GetInt16("nofont")).To.Equal(true)
  // expect(0 == decoder.GetInt32("nofont")).To.Equal(true)
  // expect(0 == decoder.GetInt64("nofont")).To.Equal(true)
  // expect(0 == decoder.GetUint("nofont")).To.Equal(true)
  // expect(0 == decoder.GetUint8("nofont")).To.Equal(true)
  // expect(0 == decoder.GetUint16("nofont")).To.Equal(true)
  // expect(0 == decoder.GetUint32("nofont")).To.Equal(true)
  // expect(0 == decoder.GetUint64("nofont")).To.Equal(true)
  // expect(0 == decoder.GetFloat32("nofont")).To.Equal(true)
  // expect(0 == decoder.GetFloat64("nofont")).To.Equal(true)
}
