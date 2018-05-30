package checker

import (
	"github.com/a8m/expect"
	"testing"
	"fmt"
)

func TestSpecChecker(t *testing.T) {
	fmt.Println()
	expect := expect.New(t)
	checker := &Checker{}
	var pint *int = nil
	expect(checker.Ensure(true)).To.Equal(checker)
	expect(checker.NotNull("str")).To.Equal(checker)
	expect(checker.NotNull(0)).To.Equal(checker)
	expect(checker.NotNull(false)).To.Equal(checker)
	expect(checker.Exec(func() { checker.NotNull(nil) }) != nil).To.Equal(true)
	expect(checker.Exec(func() { checker.NotNull(pint) }) != nil).To.Equal(true)
	n := 10
	pint = &n
	expect(checker.NotNull(pint)).To.Equal(checker)
}

func TestLengthChecker(t *testing.T) {
	expect := expect.New(t)
	checker := &Checker{}

	str := "abcd"
	expect(checker.Lgt(len(str), 2)).To.Equal(checker)
	expect(checker.Lgte(len(str), 2)).To.Equal(checker)
	expect(checker.Lgte(len(str), len(str))).To.Equal(checker)
	expect(checker.Llt(len(str), 8)).To.Equal(checker)
	expect(checker.Llte(len(str), 8)).To.Equal(checker)
	expect(checker.Llte(len(str), len(str))).To.Equal(checker)

	expect(checker.Exec(func() { checker.Lgt(len(str), 100) }) != nil).To.Equal(true)
	expect(checker.Exec(func() { checker.Lgte(len(str), 100) }) != nil).To.Equal(true)
	expect(checker.Exec(func() { checker.Llt(len(str), 0) }) != nil).To.Equal(true)
	expect(checker.Exec(func() { checker.Llte(len(str), 0) }) != nil).To.Equal(true)

	arr := []int{1, 2, 3, 4}
	expect(checker.Lgt(len(arr), 2)).To.Equal(checker)
	expect(checker.Lgte(len(arr), 2)).To.Equal(checker)
	expect(checker.Lgte(len(arr), len(arr))).To.Equal(checker)
	expect(checker.Llt(len(arr), 8)).To.Equal(checker)
	expect(checker.Llte(len(arr), 8)).To.Equal(checker)
	expect(checker.Llte(len(arr), len(arr))).To.Equal(checker)

	expect(checker.Exec(func() { checker.Lgt(len(arr), 100) }) != nil).To.Equal(true)
	expect(checker.Exec(func() { checker.Lgte(len(arr), 100) }) != nil).To.Equal(true)
	expect(checker.Exec(func() { checker.Llt(len(arr), 0) }) != nil).To.Equal(true)
	expect(checker.Exec(func() { checker.Llte(len(arr), 0) }) != nil).To.Equal(true)
}
