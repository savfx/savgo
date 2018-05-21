package query

import (
	"github.com/a8m/expect"
	"testing"
)

func TestEncode(t *testing.T) {
	expect := expect.New(t)
	expect(Encode(map[string]interface{}{})).To.Equal("")
	expect(Encode(nil)).To.Equal("")

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
