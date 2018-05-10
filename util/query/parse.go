package query

import (
  "net/url"
  "regexp"
  "fmt"
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
      // a[b][c]=d => [a, b, c]
      // a[][c]=d => [a, , c]
      // a[b][]=d => [a, b, ]
      for _, arr := range keys { // a[b][c]=d [1 4 2 3] [4 7 5 6]
        start := arr[0]
        if pos < start {
          key:= name[pos: start]
          names = append(names, key)
        }
        names = append(names, name[arr[2] : arr[3]])
        pos = arr[1]
      }
      fmt.Println(names, len(names))
    }
    // a=b a
    // a=b&a=b a[]=b&a[]=b
    // a[]=b a[]
    // a[b][]=c a[b][]
    // a[][b][c][]=d a[][b][c][]
    // fmt.Println(names, len(name))
  }
  return res, nil
}

var tokenRegexp = regexp.MustCompile(`(?:^|&)((?:[^&=])*)=?([^&]*)`)
var subRegexp = regexp.MustCompile(`\[([^\]]*)]`)
