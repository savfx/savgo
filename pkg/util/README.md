# sav-util
sav util for go


### 密码散列

- `passwd.Create`
  - 类似php的`password_hash`
  - 目前只提供bcrypt算法
- `passwd.Verify`
  - 类似php的`password_verify`
- `passwd.IsCurrent`
  - 类似php的`password_needs_rehash`
  - 由于只有一个bcrypt所以暂时没用

```go

package main

import (
  "github.com/savgo/sav-util/passwd"
  "fmt"
)

func main() {
  plaintext := "password"
  pwd, _ := passwd.Create(plaintext)
  isValid, _ := passwd.Verify(plaintext, pwd)
  fmt.Println(pwd) // $2a$08$...
  fmt.Println(len(pwd)) // 60
  fmt.Println(isValid) // true
}

```
