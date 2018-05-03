package strcase

import (
  "testing"
  "github.com/a8m/expect"
)

func TestConvert(t * testing.T) {
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

func TestConvertStringType(t * testing.T) {
  expect := expect.New(t)
  maps := map[string]string{
    "camel": "helloWorld",
    "snake": "hello_world",
    "pascal": "HelloWorld",
    "hyphen": "hello-world",
  }

  for caseType, value := range maps {
    for _, current := range maps {
      expect(ConvertStringType(current, caseType)).To.Equal(value)
    }
  }
}
