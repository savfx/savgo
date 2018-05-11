package query

import (
  "net/url"
  "regexp"
  "fmt"
  "strconv"
  "strings"
)

func Parse(path string) (map[string]interface{}, error) {
  res := make(map[string]interface{}, 0)
  parts := tokenRegexp.FindAllStringSubmatch(path, -1)
  for i := 0; i < len(parts); i++ {
    mat := parts[i]
    name, err := url.QueryUnescape(mat[1])
    if err != nil {
      return res, err
    }
    value, er2 := url.QueryUnescape(mat[2])
    if er2 != nil {
      return res, er2
    }
    if strings.Index(name, "]") == -1 || strings.Index(name, "[") == -1 {
      // a[=c a[b=c a]=c a]b=c
      if ref, ok := res[name]; ok { // a=b&a=c
        switch val := ref.(type) {
          case []interface{}:
            res[name] = append(val, value)
          default:
            tmp := append(make([]interface{}, 0), ref)
            res[name] = append(tmp, value)
        }
      } else { // a=b
        res[name] = value
      }
    } else {
      keys := subRegexp.FindAllSubmatchIndex([]byte(name), -1)
      pos := 0
      names := make([]string, 0)
      var prev string = ""
      var prevRef interface{} = res
      var currRef interface{} = nil
      // a[b][c]=d => [a, b, c]
      // a[][c]=d => [a, , c]
      // a[1][c]=d => [a, 1, c]
      // a[b][]=d => [a, b, ]
      // a[b][1]=d => [a, b, 1]
      length := len(keys) - 1
      for index, arr := range keys { // a[b][c]=d [1 4 2 3] [4 7 5 6]
        start := arr[0]
        if pos < start {
          key := name[pos: start]
          prev = key
          names = append(names, key)
        }
        key := name[arr[2] : arr[3]]
        if _, err := strconv.ParseInt(key, 10, 0); err != nil && key != "" {
          switch val := prevRef.(type) {
          case map[string]interface{}:
            if _, ok := val[prev]; !ok {
              val[prev] = make(map[string]interface{}, 0)
            }
            currRef = val[prev]
          }
        } else {
          switch val := prevRef.(type) {
          case map[string]interface{}:
            currRef := make([]interface{}, 0)
            val[prev] = currRef
          }
        }
        names = append(names, key)
        pos = arr[1]
        if length == index {
          if _, err := strconv.ParseInt(key, 10, 0); err != nil && key != "" {
            switch val := currRef.(type) {
            case map[string]interface{}:
              val[key] = value
            case []interface{}:
              val = append(val, value)
            }
          } else {
            // prevRef[prev] = append(prevRef[prev], value)
          }
        } else {
          prevRef = currRef
          prev = key
        }
      }
      fmt.Println(names)
      // reduce(names, res, func (ret interface{}, key string, next * string) interface{} {
      //   if next != nil {
      //     switch val := ret.(type) {
      //       case map[string]interface{}:
      //         if _, ok := val[key]; !ok {
      //           if _, err := strconv.ParseInt(*next, 10, 0); err != nil && *next != "" {
      //             val[key] = make(map[string]interface{}, 0)
      //           }
      //         }
      //         return val[key]
      //       // case []interface{}:
      //     }
      //     // a[]=1&a[1]=2 a[b] = 2
      //     // if _, err := strconv.ParseInt(*next, 10, 0); err != nil && *next != "" {
      //     //   switch val := ret.(type) {
      //     //   case map[string]interface{}:
      //     //     if _, ok := val[key]; !ok {
      //     //       val[key] = make(map[string]interface{}, 0)
      //     //     }
      //     //     ref = val[key]
      //     //   case []interface{}:
      //     //     tmp := make(map[string]interface{}, 0)
      //     //     val = append(val, tmp)
      //     //     ref = tmp
      //     //   }
      //     // }
      //   } else {
      //     switch val := ret.(type) {
      //     case map[string]interface{}:
      //       val[key] = value
      //     case []interface{}:
      //       val = append(val, value)
      //     }
      //     return nil
      //   }
      //   fmt.Println(key, next)
      //   return nil
      // })
    }
    //   keys := subRegexp.FindAllSubmatchIndex([]byte(name), -1)
    //   pos := 0
    //   names := make([]string, 0)
    //   // a[b][c]=d => [a, b, c]
    //   // a[][c]=d => [a, , c]
    //   // a[b][]=d => [a, b, ]
    //   for _, arr := range keys { // a[b][c]=d [1 4 2 3] [4 7 5 6]
    //     start := arr[0]
    //     if pos < start {
    //       key := name[pos: start]
    //       names = append(names, key)
    //     }
    //     key := name[arr[2] : arr[3]]
    //     names = append(names, key)
    //     pos = arr[1]
    //   }
    //   // var ref interface{} = res
    //   // length := len(names) -1
    //   // for index, key := range names {
    //   //   if index < length {
    //   //     if _, err := strconv.ParseInt(key, 10, 0); err != nil && key != "" {
    //   //       switch val := ref.(type) {
    //   //       case map[string]interface{}:
    //   //         if _, ok := val[key]; !ok {
    //   //           val[key] = make(map[string]interface{}, 0)
    //   //         }
    //   //         ref = val[key]
    //   //       case []interface{}:
    //   //         tmp := make(map[string]interface{}, 0)
    //   //         val = append(val, tmp)
    //   //         ref = tmp
    //   //       }
    //   //     }
    //   //   } else {
    //   //     if _, err := strconv.ParseInt(key, 10, 0); err != nil && key != "" {
    //   //       switch val := ref.(type) {
    //   //       case map[string]interface{}:
    //   //         val[key] = value
    //   //       }
    //   //     } else {
    //   //       switch val := ref.(type) {
    //   //       case []interface{}:
    //   //         val = append(val, value)
    //   //       }
    //   //     }
    //   //   }
    //   // }
    //   fmt.Println(names, len(names))
    // }
    // a=b a
    // a=b&a=b a[]=b&a[]=b
    // a[]=b a[]
    // a[b][]=c a[b][]
    // a[][b][c][]=d a[][b][c][]
    // fmt.Println(names, len(name))
  }
  return res, nil
}

func reduce(values []string, identity interface{}, 
  reducer func (a interface{}, b string, c * string) (interface{})) (interface{}) {
  res := identity
  length := len(values) - 1 
  for index, v := range values {
    if index < length {
      res = reducer(res, v, &values[index+1])
    } else {
      res = reducer(res, v, nil)
    }
  }
  return res
}

var tokenRegexp = regexp.MustCompile(`(?:^|&)((?:[^&=])*)=?([^&]*)`)
var subRegexp = regexp.MustCompile(`\[([^\]]*)]`)
