package convert

import (
	"github.com/a8m/expect"
	"testing"
)

func TestString(t *testing.T) {
	expect := expect.New(t)
	expect(*StringPtr("str")).To.Equal("str")
	expect(*StringPtr(`str`)).To.Equal("str")
	expect(*StringPtr([]byte(`str`))).To.Equal("str")
	expect(*StringPtr(-1)).To.Equal("-1")
	expect(*StringPtr(2)).To.Equal("2")
	expect(*StringPtr(3.14)).To.Equal("3.14")
	expect(*StringPtr(3888.1455555)).To.Equal("3888.1455555")
	expect(*StringPtr(38881455555.555)).To.Equal("38881455555.555")
	expect(*StringPtr(true)).To.Equal("true")
	expect(*StringPtr(false)).To.Equal("false")
	expect(nil == StringPtr(nil)).To.Equal(true)
}

func TestBool(t *testing.T) {
	expect := expect.New(t)
	expect(*BoolPtr("true")).To.Equal(true)
	expect(*BoolPtr("false")).To.Equal(false)
	expect(nil == BoolPtr("null")).To.Equal(true)
	expect(*BoolPtr("x")).To.Equal(false)
	expect(*BoolPtr(123)).To.Equal(true)
	expect(*BoolPtr(123.4)).To.Equal(true)
	expect(*BoolPtr(0)).To.Equal(false)
	expect(*BoolPtr(0.0000000000000000000001)).To.Equal(true)
	expect(*BoolPtr(0.0000000000000000000000)).To.Equal(false)
	expect(nil == BoolPtr(nil)).To.Equal(true)
}

func TestInt(t *testing.T) {
	expect := expect.New(t)
	expect(nil == Int64Ptr("str")).To.Equal(true)
	expect(nil == Int64Ptr("123.4")).To.Equal(true)
	expect(nil == Uint64Ptr("123.4")).To.Equal(true)
	expect(nil == Uint64Ptr("str")).To.Equal(true)

	expect(-123 == *Int64Ptr("-123")).To.Equal(true)
	expect(123 == *Int64Ptr("0123")).To.Equal(true)
	expect(123456798 == *Int64Ptr("123456798")).To.Equal(true)

	expect(nil == Int8Ptr("123456")).To.Equal(true)
	expect(nil == Int8Ptr(123456)).To.Equal(true)
	expect(-1 == *Int8Ptr("-1")).To.Equal(true)
	expect(0 == *Int8Ptr("0")).To.Equal(true)
	expect(1 == *Int8Ptr("1")).To.Equal(true)
	expect(nil == Int8Ptr(-123456)).To.Equal(true)
	expect(nil == Int8Ptr("-123456")).To.Equal(true)

	expect(nil == Uint8Ptr("123456")).To.Equal(true)
	expect(0 == *Uint8Ptr("0")).To.Equal(true)
	expect(1 == *Uint8Ptr("1")).To.Equal(true)
	expect(nil == Uint8Ptr("-1")).To.Equal(true)

	expect(nil == Int16Ptr("123456798")).To.Equal(true)
	expect(nil == Int16Ptr(123456789)).To.Equal(true)
	expect(-1 == *Int16Ptr("-1")).To.Equal(true)
	expect(0 == *Int16Ptr("0")).To.Equal(true)
	expect(1 == *Int16Ptr("1")).To.Equal(true)
	expect(nil == Int16Ptr(-123456789)).To.Equal(true)
	expect(nil == Int16Ptr("-123456789")).To.Equal(true)

	expect(nil == Uint16Ptr("123456789")).To.Equal(true)
	expect(0 == *Uint16Ptr("0")).To.Equal(true)
	expect(1 == *Uint16Ptr("1")).To.Equal(true)
	expect(nil == Uint16Ptr("-1")).To.Equal(true)

	expect(nil == Int32Ptr("1234567890123456")).To.Equal(true)
	expect(nil == Int32Ptr(1234567890123456)).To.Equal(true)
	expect(-1 == *Int32Ptr("-1")).To.Equal(true)
	expect(0 == *Int32Ptr("0")).To.Equal(true)
	expect(1 == *Int32Ptr("1")).To.Equal(true)
	expect(nil == Int32Ptr(-1234567890123456)).To.Equal(true)
	expect(nil == Int32Ptr("-1234567890123456")).To.Equal(true)

	expect(nil == Uint32Ptr("1234567890123456")).To.Equal(true)
	expect(0 == *Uint32Ptr("0")).To.Equal(true)
	expect(1 == *Uint32Ptr("1")).To.Equal(true)
	expect(nil == Uint32Ptr("-1")).To.Equal(true)
}

func TestFloat(t *testing.T) {
	expect := expect.New(t)
	expect(nil == Float64Ptr("true")).To.Equal(true)
	expect(nil == Float32Ptr("true")).To.Equal(true)
	expect(123 == *Float64Ptr("123")).To.Equal(true)
	expect(123 == *Float32Ptr("123")).To.Equal(true)
	expect(123.4 == *Float64Ptr("123.4")).To.Equal(true)
	expect(123.4 == *Float32Ptr("123.4")).To.Equal(true)
	expect(0 == *Float64Ptr("0")).To.Equal(true)
	expect(0 == *Float32Ptr("0")).To.Equal(true)
	expect(0 == *Float64Ptr(0)).To.Equal(true)
	expect(0 == *Float32Ptr(0)).To.Equal(true)
	expect(123 == *Float64Ptr(123)).To.Equal(true)
	expect(123 == *Float32Ptr(123)).To.Equal(true)
	expect(52 == *Float64Ptr(uint8(52))).To.Equal(true)
	expect(52 == *Float32Ptr(uint8(52))).To.Equal(true)
	expect(0 == *Float64Ptr("0.0000000000000000000001")).To.Equal(false)
	expect(0.0000000000000000000001 == *Float64Ptr("0.0000000000000000000001")).To.Equal(true)
	expect(0 == *Float64Ptr("0.0000000000000000000000")).To.Equal(true)

}
