package convert

import (
	"encoding/json"
	"fmt"
	"github.com/a8m/expect"
	"github.com/savfx/savgo/util/query"
	"net/url"
	"testing"
)

func TestForm(t *testing.T) {
	fmt.Println()
	expect := expect.New(t)
	text := `{
	  "name": "jetiny",
	  "age": "35",
	  "man": "true",
	  "profile": {
	     "company": "hfjy",
	     "sex": "female",
	      "xxx": {
	      	"bb": "cc"
	      },
	      "hh" : [
	      	"xx",
	      	"yy"
	      ],
	      "mm": [
	     	{"name": "张三", "age": 10}
	      ],
	      "ages": [1, 2, 3]
	  },
	  "followers": [
	    {"name": "张三", "age": 10},
	    {"name": "李四", "age": 10},
	    {"name": "王五", "age": 10}
	  ]
	}`
	jsonObject := map[string]interface{}{}
	json.Unmarshal([]byte(text), &jsonObject)
	qtext := query.Encode(jsonObject)
	values, _ := url.ParseQuery(qtext)
	root := NewFormObject(values)
	expect(root.GetString("name")).To.Equal("jetiny")
	expect(root.GetArray("followers").Len()).To.Equal(3)
	expect(root.GetInt("age")).To.Equal(35)
	expect(root.GetObject("profile").GetArray("ages").Len()).To.Equal(3)
	root.GetArray("followers").EachField(func(index int, val *FormObject) {
		expect(val.GetInt("age")).To.Equal(10)
	})
	root.GetObject("profile").GetArray("ages").EachValue(func(index int, val *ValueAccess) {
		expect(val.Int() > 0).To.Equal(true)
	})
}
