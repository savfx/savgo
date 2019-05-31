package convert

import (
	"encoding/json"
	"github.com/a8m/expect"
	"testing"
)

func TestObject(t *testing.T) {
	expect := expect.New(t)
	text := `{
    "name": "jetiny",
    "age": "35",
    "man": "true",
    "profile": {
       "company": "test"
    },
    "followers": [
      {"name": "张三", "age": 10},
      {"name": "李四", "age": 20}
    ]
  }`
	jsonText := map[string]interface{}{}
	json.Unmarshal([]byte(text), &jsonText)
	object := NewObjectAccess(jsonText)
	expect(*object.GetStringPtr("name")).To.Equal("jetiny")
	expect(object.GetString("name")).To.Equal("jetiny")
	expect(*object.GetIntPtr("age")).To.Equal(35)
	expect(object.GetInt("age")).To.Equal(35)
	expect(*object.GetBoolPtr("man")).To.Equal(true)
	expect(object.GetBool("man")).To.Equal(true)
	expect(35 == *object.GetInt8Ptr("age")).To.Equal(true)
	expect(35 == object.GetInt8("age")).To.Equal(true)
	expect(35 == *object.GetUint8Ptr("age")).To.Equal(true)
	expect(35 == object.GetUint8("age")).To.Equal(true)
	expect(35 == *object.GetInt16Ptr("age")).To.Equal(true)
	expect(35 == object.GetInt16("age")).To.Equal(true)
	expect(35 == *object.GetUint16Ptr("age")).To.Equal(true)
	expect(35 == object.GetUint16("age")).To.Equal(true)
	expect(35 == *object.GetInt32Ptr("age")).To.Equal(true)
	expect(35 == object.GetInt32("age")).To.Equal(true)
	expect(35 == *object.GetUint32Ptr("age")).To.Equal(true)
	expect(35 == object.GetUint32("age")).To.Equal(true)
	expect(35 == *object.GetInt64Ptr("age")).To.Equal(true)
	expect(35 == object.GetInt64("age")).To.Equal(true)
	expect(35 == *object.GetUint64Ptr("age")).To.Equal(true)
	expect(35 == object.GetUint64("age")).To.Equal(true)
	expect(35 == *object.GetFloat32Ptr("age")).To.Equal(true)
	expect(35 == object.GetFloat32("age")).To.Equal(true)
	expect(35 == *object.GetFloat64Ptr("age")).To.Equal(true)
	expect(35 == object.GetFloat64("age")).To.Equal(true)

	expect(object.GetString("nofont")).To.Equal("")
	expect(false == object.GetBool("nofont")).To.Equal(true)
	expect(0 == object.GetInt("nofont")).To.Equal(true)
	expect(0 == object.GetInt8("nofont")).To.Equal(true)
	expect(0 == object.GetInt16("nofont")).To.Equal(true)
	expect(0 == object.GetInt32("nofont")).To.Equal(true)
	expect(0 == object.GetInt64("nofont")).To.Equal(true)
	expect(0 == object.GetUint("nofont")).To.Equal(true)
	expect(0 == object.GetUint8("nofont")).To.Equal(true)
	expect(0 == object.GetUint16("nofont")).To.Equal(true)
	expect(0 == object.GetUint32("nofont")).To.Equal(true)
	expect(0 == object.GetUint64("nofont")).To.Equal(true)
	expect(0 == object.GetFloat32("nofont")).To.Equal(true)
	expect(0 == object.GetFloat64("nofont")).To.Equal(true)

	expect(object.Has("profile")).To.Equal(true)
	expect(object.HasObject("profile")).To.Equal(true)
	expect(object.GetObject("profile").GetString("company")).To.Equal("hfjy")
	expect(object.Has("nofound")).To.Equal(false)
	expect(object.HasObject("nofound")).To.Equal(false)

	expect(object.Has("followers")).To.Equal(true)
	expect(object.HasArray("followers")).To.Equal(true)
	expect(object.GetArray("followers").Len()).To.Equal(2)
	expect(object.GetArray("followers").Field(0).Object().GetString("name")).To.Equal("张三")
	expect(object.GetValue("followers").IsArray()).To.Equal(true)

	object.ForEach(func(name string, value *ValueAccess) bool {
		return false
	})
}
