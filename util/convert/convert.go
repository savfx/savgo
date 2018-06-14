package convert

import (
	"strconv"
)

func StringPtr(value interface{}) *string {
	var res string
	switch val := value.(type) {
	case string:
		res = val
	case bool:
		if val {
			res = "true"
		} else {
			res = "false"
		}
	case int:
		res = strconv.FormatInt(int64(val), 10)
	case uint:
		res = strconv.FormatUint(uint64(val), 10)
	case int8:
		res = strconv.FormatInt(int64(val), 10)
	case uint8:
		res = strconv.FormatUint(uint64(val), 10)
	case int16:
		res = strconv.FormatInt(int64(val), 10)
	case uint16:
		res = strconv.FormatUint(uint64(val), 10)
	case int32:
		res = strconv.FormatInt(int64(val), 10)
	case uint32:
		res = strconv.FormatUint(uint64(val), 10)
	case int64:
		res = strconv.FormatInt(val, 10)
	case uint64:
		res = strconv.FormatUint(val, 10)
	case float32:
		res = strconv.FormatFloat(float64(val), 'f', -1, 64)
	case float64:
		res = strconv.FormatFloat(val, 'f', -1, 64)
	case []byte:
		res = string(val)
	default:
		return nil
	}
	return &res
}

func BoolPtr(value interface{}) *bool {
	var res bool
	switch val := value.(type) {
	case bool:
		res = val
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		res = val != 0
	case float32:
		if val != 0 {
			res = true
		} else {
			res = false
		}
	case float64:
		if val != 0 {
			res = true
		} else {
			res = false
		}
	case string:
		switch val {
		case "1", "true", "TRUE", "True", "On":
			res = true
		case "null":
			return nil
		default:
			res = false
		}
	case []byte:
		switch string(val) {
		case "1", "true", "TRUE", "True", "On":
			res = true
		case "null":
			return nil
		default:
			res = false
		}
	default:
		return nil
	}
	return &res
}

// int
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
const INT_MAX_64 = int64(INT_MAX)
const INT_MIN_64 = int64(INT_MIN)

// uint
const UINT_MIN uint = 0
const UINT_MAX = ^uint(0)
const UINT_MAX_64 = uint64(UINT_MAX)
const UINT_MIN_64 = uint64(UINT_MIN)

// int8 -127 ~ 128
const INT8_MAX = int8(^uint8(0) >> 1)
const INT8_MIN = ^INT8_MAX
const INT8_MAX_64 = int64(INT8_MAX)
const INT8_MIN_64 = int64(INT8_MIN)

// uint8 0 ~ 255
const UINT8_MAX = ^uint8(0)
const UINT8_MIN = uint8(0)
const UINT8_MAX_64 = uint64(UINT8_MAX)
const UINT8_MIN_64 = uint64(UINT8_MIN)

// int16 -32768 ~ 32767
const INT16_MAX = int16(^uint16(0) >> 1)
const INT16_MIN = ^INT16_MAX
const INT16_MAX_64 = int64(INT16_MAX)
const INT16_MIN_64 = int64(INT16_MIN)

// uint16 0 ~ 65535
const UINT16_MAX = ^uint16(0)
const UINT16_MIN = uint16(0)
const UINT16_MAX_64 = uint64(UINT16_MAX)
const UINT16_MIN_64 = uint64(UINT16_MIN)

// int32 -2147483648 ~ 2147483647
const INT32_MAX = int32(^uint32(0) >> 1)
const INT32_MIN = ^INT32_MAX
const INT32_MAX_64 = int64(INT32_MAX)
const INT32_MIN_64 = int64(INT32_MIN)

// uint32 0 ~ 4294967295
const UINT32_MAX = ^uint32(0)
const UINT32_MIN = uint32(0)
const UINT32_MAX_64 = uint64(UINT32_MAX)
const UINT32_MIN_64 = uint64(UINT32_MIN)

func IntPtr(value interface{}) *int {
	ptr := Int64Ptr(value)
	if ptr == nil || *ptr > INT_MAX_64 || *ptr < INT_MIN_64 {
		return nil
	}
	val := int(*ptr)
	return &val
}

func UintPtr(value interface{}) *uint {
	ptr := Uint64Ptr(value)
	if ptr == nil || *ptr > UINT_MAX_64 || *ptr < UINT_MIN_64 {
		return nil
	}
	val := uint(*ptr)
	return &val
}

func Int8Ptr(value interface{}) *int8 {
	ptr := Int64Ptr(value)
	if ptr == nil || *ptr > INT8_MAX_64 || *ptr < INT8_MIN_64 {
		return nil
	}
	val := int8(*ptr)
	return &val
}

func Uint8Ptr(value interface{}) *uint8 {
	ptr := Uint64Ptr(value)
	if ptr == nil || *ptr > UINT8_MAX_64 || *ptr < UINT8_MIN_64 {
		return nil
	}
	val := uint8(*ptr)
	return &val
}

func Int16Ptr(value interface{}) *int16 {
	ptr := Int64Ptr(value)
	if ptr == nil || *ptr > INT16_MAX_64 || *ptr < INT16_MIN_64 {
		return nil
	}
	val := int16(*ptr)
	return &val
}

func Uint16Ptr(value interface{}) *uint16 {
	ptr := Uint64Ptr(value)
	if ptr == nil || *ptr > UINT16_MAX_64 || *ptr < UINT16_MIN_64 {
		return nil
	}
	val := uint16(*ptr)
	return &val
}

func Int32Ptr(value interface{}) *int32 {
	ptr := Int64Ptr(value)
	if ptr == nil || *ptr > INT32_MAX_64 || *ptr < INT32_MIN_64 {
		return nil
	}
	val := int32(*ptr)
	return &val
}

func Uint32Ptr(value interface{}) *uint32 {
	ptr := Uint64Ptr(value)
	if ptr == nil || *ptr > UINT32_MAX_64 || *ptr < UINT32_MIN_64 {
		return nil
	}
	val := uint32(*ptr)
	return &val
}

func Int64Ptr(value interface{}) *int64 {
	var res int64
	switch val := value.(type) {
	case bool:
		if val {
			res = 1
		} else {
			res = 0
		}
	case int:
		res = int64(val)
	case uint:
		res = int64(val)
	case int8:
		res = int64(val)
	case uint8:
		res = int64(val)
	case int16:
		res = int64(val)
	case uint16:
		res = int64(val)
	case int32:
		res = int64(val)
	case uint32:
		res = int64(val)
	case int64:
		res = val
	case uint64:
		res = int64(val)
	case float32:
		res = int64(val)
	case float64:
		res = int64(val)
	case []byte:
		if ret, err := strconv.ParseInt(string(val), 10, 0); err == nil {
			res = ret
		} else {
			return nil
		}
	case string:
		if ret, err := strconv.ParseInt(val, 10, 0); err == nil {
			res = ret
		} else {
			return nil
		}
	default:
		return nil
	}
	return &res
}

func Uint64Ptr(value interface{}) *uint64 {
	switch val := value.(type) {
	case uint64:
		res := val
		return &res
	case []byte:
		if ret, err := strconv.ParseUint(string(val), 10, 0); err == nil {
			return &ret
		}
	case string:
		if ret, err := strconv.ParseUint(val, 10, 0); err == nil {
			return &ret
		}
	default:
		intVal := Int64Ptr(value)
		if intVal != nil {
			res := uint64(*intVal)
			return &res
		}
	}
	return nil
}

func Float32Ptr(value interface{}) *float32 {
	ptr := Float64Ptr(value)
	if ptr == nil {
		return nil
	}
	val := float32(*ptr)
	return &val
}

func Float64Ptr(value interface{}) *float64 {
	var res float64
	switch val := value.(type) {
	case bool:
		if val {
			res = 1
		} else {
			res = 0
		}
	case int:
		res = float64(val)
	case uint:
		res = float64(val)
	case int8:
		res = float64(val)
	case uint8:
		res = float64(val)
	case int16:
		res = float64(val)
	case uint16:
		res = float64(val)
	case int32:
		res = float64(val)
	case uint32:
		res = float64(val)
	case int64:
		res = float64(val)
	case uint64:
		res = float64(val)
	case float32:
		res = float64(val)
	case float64:
		res = val
	case []byte:
		if ret, err := strconv.ParseFloat(string(val), 10); err == nil {
			res = ret
		} else {
			return nil
		}
	case string:
		if ret, err := strconv.ParseFloat(val, 10); err == nil {
			res = ret
		} else {
			return nil
		}
	default:
		return nil
	}
	return &res
}

func StringVal(value interface{}) (res string) {
	if ptr := StringPtr(value); ptr != nil {
		res = *ptr
	}
	return
}

func BoolVal(value interface{}) (res bool) {
	if ptr := BoolPtr(value); ptr != nil {
		res = *ptr
	}
	return
}

func IntVal(value interface{}) (res int) {
	if ptr := IntPtr(value); ptr != nil {
		res = *ptr
	}
	return
}

func UintVal(value interface{}) (res uint) {
	if ptr := UintPtr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Int8Val(value interface{}) (res int8) {
	if ptr := Int8Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Uint8Val(value interface{}) (res uint8) {
	if ptr := Uint8Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Int16Val(value interface{}) (res int16) {
	if ptr := Int16Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Uint16Val(value interface{}) (res uint16) {
	if ptr := Uint16Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Int32Val(value interface{}) (res int32) {
	if ptr := Int32Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Uint32Val(value interface{}) (res uint32) {
	if ptr := Uint32Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Int64Val(value interface{}) (res int64) {
	if ptr := Int64Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Uint64Val(value interface{}) (res uint64) {
	if ptr := Uint64Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Float32Val(value interface{}) (res float32) {
	if ptr := Float32Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}

func Float64Val(value interface{}) (res float64) {
	if ptr := Float64Ptr(value); ptr != nil {
		res = *ptr
	}
	return
}
