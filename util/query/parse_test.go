package query

import (
  "github.com/a8m/expect"
  "testing"
  "fmt"
)

func TestParse(t *testing.T) {
  expect := expect.New(t)
  // m, _ := Parse("a+=b&c=d&c=e&x[bbc=2&x[bbc=3")
  // expect(len(m) >=0).To.Equal(true)
  // fmt.Println(m, len(m))
  // m, _ := Parse("a[b][c]=d&a[m][c]=d&a[b][d]=d&x[]=y")
  // m, _ := Parse("a[b][c]=d&a[b][f]=e")
  m, _ := Parse("a[][d]=d")
  expect(len(m) >=0).To.Equal(true)
  fmt.Println(m, len(m))
}
