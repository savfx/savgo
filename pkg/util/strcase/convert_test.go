package strcase

import (
  "testing"
  "github.com/a8m/expect"
)

func TestCamel(t * testing.T) {
  expect := expect.New(t)
  maps := map[CaseType]string{
    CamelType: "helloWorld",
    SnakeType: "hello_world",
    PascalType: "HelloWorld",
    HyphenType: "hello-world",
  }

  for caseType, value := range maps {
    for _, current := range maps {
      expect(Convert(current, caseType)).To.Equal(value)
    }
  }
}
