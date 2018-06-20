package query

import (
	"net/url"
	"strconv"
	"strings"
)

func Encode(value map[string]interface{}) string {
	arr := make([]string, 0)
	if value != nil {
		for name, val := range value {
			buildParams(url.QueryEscape(name), val, &arr)
		}
	}
	return strings.Join(arr, "&")
}

func appendValue(prefix string, value interface{}, arr *[]string) {
	str := ""
	switch val := value.(type) {
	case string:
		str = url.QueryEscape(val)
	case *string:
		str = url.QueryEscape(*val)
	case []byte:
		str = url.QueryEscape(string(val))
	case *[]byte:
		str = url.QueryEscape(string(*val))
	case bool:
		if val {
			str = "true"
		} else {
			str = "false"
		}
	case *bool:
		if *val {
			str = "true"
		} else {
			str = "false"
		}
	case int:
		str = strconv.FormatInt(int64(val), 10)
	case *int:
		str = strconv.FormatInt(int64(*val), 10)
	case int8:
		str = strconv.FormatInt(int64(val), 10)
	case *int8:
		str = strconv.FormatInt(int64(*val), 10)
	case int16:
		str = strconv.FormatInt(int64(val), 10)
	case int32:
		str = strconv.FormatInt(int64(val), 10)
	case int64:
		str = strconv.FormatInt(val, 10)
	case uint:
		str = strconv.FormatUint(uint64(val), 10)
	case uint8:
		str = strconv.FormatUint(uint64(val), 10)
	case uint16:
		str = strconv.FormatUint(uint64(val), 10)
	case uint32:
		str = strconv.FormatUint(uint64(val), 10)
	case uint64:
		str = strconv.FormatUint(val, 10)
	case float32:
		str = strconv.FormatFloat(float64(val), 'f', -1, 64)
	case float64:
		str = strconv.FormatFloat(val, 'f', -1, 64)
	case *int16:
		str = strconv.FormatInt(int64(*val), 10)
	case *int32:
		str = strconv.FormatInt(int64(*val), 10)
	case *int64:
		str = strconv.FormatInt(*val, 10)
	case *uint:
		str = strconv.FormatUint(uint64(*val), 10)
	case *uint8:
		str = strconv.FormatUint(uint64(*val), 10)
	case *uint16:
		str = strconv.FormatUint(uint64(*val), 10)
	case *uint32:
		str = strconv.FormatUint(uint64(*val), 10)
	case *uint64:
		str = strconv.FormatUint(*val, 10)
	case *float32:
		str = strconv.FormatFloat(float64(*val), 'f', -1, 64)
	case *float64:
		str = strconv.FormatFloat(*val, 'f', -1, 64)
	}
	*arr = append(*arr, prefix+"="+str)
}

func buildParams(prefix string, value interface{}, arr *[]string) {
	switch val := value.(type) {
	case []interface{}:
		length := len(prefix)
		if length > 1 && prefix[length-2:] == "[]" {
			for _, v := range val {
				appendValue(prefix, v, arr)
			}
		} else {
			for index, v := range val {
				switch v.(type) {
				case map[string]interface{}:
					buildParams(prefix+"["+strconv.Itoa(index)+"]", v, arr)
				default:
					buildParams(prefix+"[]", v, arr)
				}
			}
		}
	case map[string]interface{}:
		if val != nil {
			for k, v := range val {
				buildParams(prefix+"["+url.QueryEscape(k)+"]", v, arr)
			}
		} else {
			appendValue(prefix, value, arr)
		}
	default:
		appendValue(prefix, value, arr)
	}
}

// func Encode(value map[string]interface{}) (string) {
// 	return buildQuery(value, "")
// }

// func buildQuery(value interface{}, prefix string) (string) {
// 	switch val := value.(type) {
// 	case []interface{}:
//     arr := make([]string, 0)
// 		for _, v := range val {
// 			part := buildQuery(v, prefix+"[]")
//       if part != "" {
//         arr = append(arr, part)
//       }
// 		}
//     return strings.Join(arr, "&")
// 	case map[string]interface{}:
// 		length := len(val)
//     arr := make([]string, 0)
// 		for k, v := range val {
//   		length -= 1
//       if k != "" {
//         childPrefix := ""
//         if prefix != "" {
//           childPrefix = prefix + "[" + url.QueryEscape(k) + "]"
//         } else {
//           childPrefix = url.QueryEscape(k)
//         }
//         part := buildQuery(v, childPrefix)
//         if part != "" {
//           arr = append(arr, part)
//         }
//       }
// 		}
//     return strings.Join(arr, "&")
// 	case string:
// 	  return prefix + "=" + url.QueryEscape(val)
//   case []byte:
//     return prefix + "=" + url.QueryEscape(string(val))
//   case bool:
//     if val {
//       return prefix + "=true"
//     } else {
//       return prefix + "=false"
//     }
//   case int:
//     return prefix + "=" + strconv.FormatInt(int64(val), 10)
//   case int8:
//     return prefix + "=" + strconv.FormatInt(int64(val), 10)
//   case int16:
//     return prefix + "=" + strconv.FormatInt(int64(val), 10)
//   case int32:
//     return prefix + "=" + strconv.FormatInt(int64(val), 10)
//   case int64:
//     return prefix + "=" + strconv.FormatInt(val, 10)
//   case uint:
//     return prefix + "=" + strconv.FormatUint(uint64(val), 10)
//   case uint8:
//     return prefix + "=" + strconv.FormatUint(uint64(val), 10)
//   case uint16:
//     return prefix + "=" + strconv.FormatUint(uint64(val), 10)
//   case uint32:
//     return prefix + "=" + strconv.FormatUint(uint64(val), 10)
//   case uint64:
//     return prefix + "=" + strconv.FormatUint(val, 10)
//   case float32:
//     return prefix + "=" + strconv.FormatFloat(float64(val), 'f', -1, 64)
//   case float64:
//     return prefix + "=" + strconv.FormatFloat(val, 'f', -1, 64)
// 	}
// 	return prefix
// }
