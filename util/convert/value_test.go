package convert

import (
	"github.com/a8m/expect"
	"testing"
)

func TestValue(t *testing.T) {
	expect := expect.New(t)
	arr := []interface{}{"string"}
	expect(NewValueAccess(arr).IsArray()).To.Equal(true)
	expect(NewValueAccess(arr).IsObject()).To.Equal(false)
	expect(NewValueAccess(arr).Array().Len()).To.Equal(1)
	obj := map[string]interface{}{}
	expect(NewValueAccess(obj).IsArray()).To.Equal(false)
	expect(NewValueAccess(obj).IsObject()).To.Equal(true)
	expect(NewValueAccess(obj).Object().GetString("nofound")).To.Equal("")

	val := NewValueAccess("str")
	expect(nil == val.StringPtr()).To.Equal(false)
	expect(nil == val.BoolPtr()).To.Equal(false)
	expect(nil == val.IntPtr()).To.Equal(true)
	expect(nil == val.Int8Ptr()).To.Equal(true)
	expect(nil == val.Int16Ptr()).To.Equal(true)
	expect(nil == val.Int32Ptr()).To.Equal(true)
	expect(nil == val.Int64Ptr()).To.Equal(true)
	expect(nil == val.UintPtr()).To.Equal(true)
	expect(nil == val.Uint8Ptr()).To.Equal(true)
	expect(nil == val.Uint16Ptr()).To.Equal(true)
	expect(nil == val.Uint32Ptr()).To.Equal(true)
	expect(nil == val.Uint64Ptr()).To.Equal(true)
	expect(nil == val.Float32Ptr()).To.Equal(true)
	expect(nil == val.Float64Ptr()).To.Equal(true)

	expect(val.String()).To.Equal("str")
	expect(val.Bool()).To.Equal(false)
	expect(0 == val.Int()).To.Equal(true)
	expect(0 == val.Int8()).To.Equal(true)
	expect(0 == val.Int16()).To.Equal(true)
	expect(0 == val.Int32()).To.Equal(true)
	expect(0 == val.Int64()).To.Equal(true)
	expect(0 == val.Uint()).To.Equal(true)
	expect(0 == val.Uint8()).To.Equal(true)
	expect(0 == val.Uint16()).To.Equal(true)
	expect(0 == val.Uint32()).To.Equal(true)
	expect(0 == val.Uint64()).To.Equal(true)
	expect(0 == val.Float32()).To.Equal(true)
	expect(0 == val.Float64()).To.Equal(true)

}
