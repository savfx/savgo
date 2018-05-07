package convert

import (
  "github.com/a8m/expect"
  "testing"
)

func TestArray(t *testing.T) {
  expect := expect.New(t)
  arr := NewArrayAccess([]interface{}{
    "jetiny",
    "35",
    "true",
    []interface{}{
      "hello",
    },
    map[string]interface{}{
      "name": "jetiny",
    },
  })
  expect(arr.Len()).To.Equal(5)
  expect(arr.Field(0).String()).To.Equal("jetiny")
  expect(arr.Field(1).Int()).To.Equal(35)
  expect(arr.Field(2).Bool()).To.Equal(true)
  expect(arr.Field(2).Array().Len()).To.Equal(0)
  expect(arr.Field(3).Array().Field(0).String()).To.Equal("hello")
  expect(arr.Field(4).Array().Len()).To.Equal(0)
  expect(arr.Field(4).Object().GetValue("name").String()).To.Equal("jetiny")

  arr.ForEach(func (index int, value * ValueAccess ) bool {
    return false
  })
}
