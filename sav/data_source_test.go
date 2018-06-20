package sav

import (
	"encoding/json"
	"github.com/a8m/expect"
	"net/url"
	"testing"
)

func TestNewJsonDataSource(t *testing.T) {
	expect := expect.New(t)
	{
		ds := NewJsonDataSource([]byte(`[{"a": 2}]`), json.Unmarshal)
		expect(ds != nil).To.Be.True()
		expect(ds.GetArrayAccess() != nil).To.Be.True()
	}
	{
		ds := NewJsonDataSource([]byte(`{"a": 2}`), json.Unmarshal)
		expect(ds != nil).To.Be.True()
		expect(ds.GetObjectAccess() != nil).To.Be.True()
	}
}

func TestNewFormDataSource(t *testing.T) {
	expect := expect.New(t)
	value := url.Values{}
	value.Add("a", "b")
	value.Add("arr[]", "c")
	value.Add("arr[]", "d")
	ds := NewFormDataSource(value)
	expect(ds != nil).To.Be.True()
	expect(ds.GetFormObject() != nil).To.Be.True()

	ds.FormArray = ds.GetFormObject().GetArray("arr")
	expect(ds.GetFormArray() != nil).To.Be.True()
	expect(ds.GetFormArray().Value(0).String()).To.Equal("c")
}
