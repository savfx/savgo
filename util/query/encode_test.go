package query

import (
	"github.com/a8m/expect"
	"testing"
	"fmt"
	"encoding/json"
)

func TestEncode(t *testing.T) {
	expect := expect.New(t)
	expect(Encode(map[string]interface{}{})).To.Equal("")
	expect(Encode(nil)).To.Equal("")
	var x string = "ss"
	var n uint8 = 20
	expect(Encode(map[string]interface{}{"x": &x, "n": &n})).To.Equal(`x=ss&n=20`)

	// expect(Encode(map[string]interface{}{
	// 	"a": []interface{}{
	// 		map[string]interface{}{
	// 			"":  "skip",
	// 			"b": []byte("1"),
	// 		},
	// 		map[string]interface{}{"": "skip"},
	// 		map[string]interface{}{"": "skip"},
	// 		map[string]interface{}{
	// 			"c": false,
	// 			"":  "skip",
	// 		},
	// 		map[string]interface{}{"": "skip"},
	// 		map[string]interface{}{
	// 			"d": true,
	// 		},
	// 		map[string]interface{}{"": "skip"},
	// 	},
	// 	"": "skip",
	// })).To.Equal("a[][b]=1&a[][c]=false&a[][d]=true")

	expect(Encode(map[string]interface{}{
		"a": []interface{}{
			int(1),
			int8(2),
			int16(3),
			int32(4),
			int64(5),
			uint(6),
			uint8(7),
			uint16(8),
			uint32(9),
			uint64(10),
			float32(11),
			float64(12),
		},
	})).To.Equal("a[]=1&a[]=2&a[]=3&a[]=4&a[]=5&a[]=6&a[]=7&a[]=8&a[]=9&a[]=10&a[]=11&a[]=12")

	expect(Encode(map[string]interface{}{
		"a": []interface{}{
			map[string]interface{}{
				"b": "c",
			},
		},
	})).To.Equal("a[0][b]=c")

}

func TestEncode2(t *testing.T) {
	expect := expect.New(t)
	text := `{
	  "name": "jetiny",
	  "ages": [1, 2, 3],
	  "followers": [
	    {"name": "san", "age": 10},
	    {"name": "si", "age": 10},
	    {"name": "wu", "age": 10}
	  ]
	}`
	jsonObject := map[string]interface{}{}
	json.Unmarshal([]byte(text), &jsonObject)
	value := Encode(jsonObject)
	fmt.Println()
	expect(value).To.Equal("name=jetiny&ages[]=1&ages[]=2&ages[]=3&followers[0][name]=san&followers[0][age]=10&followers[1][name]=si&followers[1][age]=10&followers[2][name]=wu&followers[2][age]=10")
}
